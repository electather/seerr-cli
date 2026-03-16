package mcp

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"seerr-cli/cmd/apiutil"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Default timeout values for the MCP HTTP server.
const (
	httpReadHeaderTimeout = 5 * time.Second
	httpReadTimeout       = 15 * time.Second
	httpWriteTimeout      = 30 * time.Second
	httpIdleTimeout       = 60 * time.Second
	httpShutdownTimeout   = 30 * time.Second
)

// NewHTTPServer creates an http.Server bound to addr with safe default timeouts.
// It is exported so tests can assert that the server is properly configured.
func NewHTTPServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		ReadHeaderTimeout: httpReadHeaderTimeout,
		ReadTimeout:       httpReadTimeout,
		WriteTimeout:      httpWriteTimeout,
		IdleTimeout:       httpIdleTimeout,
	}
}

var buildVersion = "dev"

// SetVersionInfo injects the linker-set build version so the MCP server can
// advertise the real application version to clients.
func SetVersionInfo(version string) {
	buildVersion = version
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MCP server",
	Long:  `Start a Model Context Protocol server that exposes the Seerr API as tools.`,
	Example: `  # Start with stdio transport (default, for Claude Desktop)
  seerr-cli mcp serve

  # Start over HTTP with Bearer token auth
  seerr-cli mcp serve --transport http --auth-token mysecret

  # Start over HTTPS with TLS
  seerr-cli mcp serve --transport http --auth-token mysecret --tls-cert /path/to/cert.pem --tls-key /path/to/key.pem

  # Accept Seerr API key via X-Api-Key header or ?api_key= query parameter
  seerr-cli mcp serve --transport http --allow-api-key-query-param

  # Enable CORS for browser-based clients (e.g. claude.ai)
  seerr-cli mcp serve --transport http --allow-api-key-query-param --cors

  # Start over HTTP without auth (insecure, not recommended)
  seerr-cli mcp serve --transport http --no-auth`,
	RunE: runServe,
}

func init() {
	RegisterFlags(serveCmd)
	BindFlags(serveCmd)
	Cmd.AddCommand(serveCmd)
}

func runServe(_ *cobra.Command, args []string) error {
	transport := viper.GetString("mcp.transport")
	addr := viper.GetString("mcp.addr")
	authToken := viper.GetString("mcp.auth_token")
	noAuth := viper.GetBool("mcp.no_auth")
	tlsCert := viper.GetString("mcp.tls_cert")
	tlsKey := viper.GetString("mcp.tls_key")
	cors := viper.GetBool("mcp.cors")
	allowAPIKeyQueryParam := viper.GetBool("mcp.allow_api_key_query_param")
	logFile := viper.GetString("mcp.log_file")
	logLevel := viper.GetString("mcp.log_level")
	logFormat := viper.GetString("mcp.log_format")

	if err := initLogger(transport, logFile, logLevel, logFormat); err != nil {
		return err
	}

	if err := ValidateServeConfig(); err != nil {
		return err
	}

	if transport == "http" && authToken == "" && !allowAPIKeyQueryParam && !noAuth {
		return fmt.Errorf("HTTP transport requires --auth-token, --allow-api-key-query-param, or --no-auth (insecure) to be set explicitly")
	}

	s := server.NewMCPServer("electather/seerr-cli", buildVersion)

	registerStatusTools(s)
	registerSearchTools(s)
	registerMoviesTools(s)
	registerTVTools(s)
	registerRequestTools(s)
	registerMediaTools(s)
	registerUsersTools(s)
	registerIssueTools(s)
	registerPersonTools(s)
	registerCollectionTools(s)
	registerServiceTools(s)
	registerSettingsTools(s)
	registerWatchlistTools(s)
	registerBlocklistTools(s)
	registerAuthTools(s)
	registerResources(s)
	registerPrompts(s)

	seerServer := viper.GetString("seerr.server")

	switch transport {
	case "stdio":
		mcpLog.Info("starting MCP server",
			"transport", "stdio",
			"seerr_api", seerServer,
			"tools", 46, "resources", 9, "prompts", 6,
		)
		mcpLog.Debug("stdio transport ready, waiting for MCP client on stdin")
		return server.ServeStdio(s)
	case "http":
		scheme := "http"
		if tlsCert != "" && tlsKey != "" {
			scheme = "https"
		}
		// Build a human-readable host for the startup message.
		host := addr
		if strings.HasPrefix(host, ":") {
			host = "localhost" + host
		}
		endpoint := fmt.Sprintf("%s://%s/mcp", scheme, host)

		mcpLog.Info("starting MCP server",
			"transport", "http",
			"endpoint", endpoint,
			"seerr_api", seerServer,
			"tools", 46, "resources", 9, "prompts", 6,
			"tls", tlsCert != "",
			"auth_token", authToken != "",
			"cors", cors,
			"allow_api_key_query_param", allowAPIKeyQueryParam,
		)

		httpHandler := server.NewStreamableHTTPServer(s)
		handler := http.Handler(httpHandler)
		// Per-request Seerr API key injection (header or optional query param).
		handler = SeerrAPIKeyMiddleware(allowAPIKeyQueryParam, handler)
		handler = httpLoggingMiddleware(handler)
		if authToken != "" {
			handler = bearerAuthMiddleware(authToken, handler)
		}
		// The health endpoint must be reachable without auth, so register it in
		// a top-level mux that sits above the bearer-auth middleware.
		topMux := http.NewServeMux()
		topMux.HandleFunc("/health", HealthCheckHandler)
		topMux.Handle("/", handler)
		handler = topMux
		// CORS must be outermost so browser preflight OPTIONS requests are never
		// blocked by auth middleware.
		if cors {
			handler = corsMiddleware(handler)
		}
		srv := NewHTTPServer(addr, handler)

		// Catch SIGINT and SIGTERM so the server shuts down gracefully.
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

		serveErrCh := make(chan error, 1)
		if tlsCert != "" && tlsKey != "" {
			go func() { serveErrCh <- srv.ListenAndServeTLS(tlsCert, tlsKey) }()
		} else {
			go func() { serveErrCh <- srv.ListenAndServe() }()
		}

		select {
		case err := <-serveErrCh:
			// Server exited on its own (e.g. port already in use).
			if err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		case <-sigCh:
			mcpLog.Info("shutting down MCP HTTP server")
		}

		shutdownCtx, cancel := context.WithTimeout(context.Background(), httpShutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("graceful shutdown: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unknown transport %q: must be stdio or http", transport)
	}
}

// ValidateServeConfig checks that the Seerr server URL is configured. It is
// exported so that tests can verify the fail-fast behaviour without starting
// the server.
func ValidateServeConfig() error {
	if apiutil.NormalizeServerURL(viper.GetString("seerr.server")) == "" {
		return fmt.Errorf("seerr.server is not configured; set it with --server <url> or add seerr.server to ~/.seerr-cli.yaml")
	}
	return nil
}

// HealthCheckHandler responds to GET /health with a JSON status payload.
// It is exported so that it can be tested directly from the tests package.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","version":%q}`, buildVersion)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Mcp-Session-Id, Accept")
		w.Header().Set("Access-Control-Expose-Headers", "Mcp-Session-Id")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func bearerAuthMiddleware(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		provided := strings.TrimPrefix(authHeader, prefix)
		if subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// SeerrAPIKeyMiddleware extracts the Seerr API key from the incoming request
// and injects it into the request context for use by MCP tool handlers.
//
// The key is read from the X-Api-Key request header first. When
// allowQueryParam is true the middleware also accepts the key via the
// api_key query parameter; the header takes precedence when both are present.
//
// If neither location provides a key the middleware responds with 401.
func SeerrAPIKeyMiddleware(allowQueryParam bool, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var apiKey string
		if v := r.Header.Get("X-Api-Key"); v != "" {
			apiKey = v
		} else if allowQueryParam {
			if v := r.URL.Query().Get("api_key"); v != "" {
				apiKey = v
			}
		}

		if apiKey != "" {
			ctx := context.WithValue(r.Context(), apiKeyCtxKey, apiKey)
			r = r.Clone(ctx)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
