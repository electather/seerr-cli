package mcp

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"strings"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

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

  # Start over HTTP with a secret path prefix (for clients that cannot send custom headers)
  seerr-cli mcp serve --transport http --route-token abc123 --no-auth
  # MCP endpoint becomes: http://localhost:8811/abc123/mcp

  # Enable CORS for browser-based clients (e.g. claude.ai)
  seerr-cli mcp serve --transport http --route-token abc123 --no-auth --cors

  # Combine both auth methods for defense in depth
  seerr-cli mcp serve --transport http --route-token abc123 --auth-token mysecret

  # Start over HTTP without auth (insecure, not recommended)
  seerr-cli mcp serve --transport http --no-auth

  # Start in multi-tenant mode (per-user API keys in URL path)
  seerr-cli mcp serve --transport http --no-auth --multi-tenant`,
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
	routeToken := viper.GetString("mcp.route_token")
	noAuth := viper.GetBool("mcp.no_auth")
	tlsCert := viper.GetString("mcp.tls_cert")
	tlsKey := viper.GetString("mcp.tls_key")
	cors := viper.GetBool("mcp.cors")
	multiTenant := viper.GetBool("mcp.multi_tenant")
	logFile := viper.GetString("mcp.log_file")
	logLevel := viper.GetString("mcp.log_level")
	logFormat := viper.GetString("mcp.log_format")

	if err := initLogger(transport, logFile, logLevel, logFormat); err != nil {
		return err
	}

	if transport == "http" && authToken == "" && routeToken == "" && !noAuth {
		return fmt.Errorf("HTTP transport requires --auth-token, --route-token, or --no-auth (insecure) to be set explicitly")
	}

	if multiTenant && transport != "http" {
		return fmt.Errorf("--multi-tenant requires --transport http")
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
		mcpPath := "/mcp"
		if multiTenant {
			mcpPath = "/{seerr-api-token}/mcp"
		} else if routeToken != "" {
			mcpPath = "/" + routeToken + "/mcp"
		}
		endpoint := fmt.Sprintf("%s://%s%s", scheme, host, mcpPath)

		mcpLog.Info("starting MCP server",
			"transport", "http",
			"endpoint", endpoint,
			"seerr_api", seerServer,
			"tools", 46, "resources", 9, "prompts", 6,
			"tls", tlsCert != "",
			"auth_token", authToken != "",
			"route_token", routeToken != "",
			"cors", cors,
			"multi_tenant", multiTenant,
		)

		httpHandler := server.NewStreamableHTTPServer(s)
		var handler http.Handler
		if multiTenant {
			handler = tenantRoutingHandler(httpHandler)
		} else if routeToken != "" {
			// Strip the route-token prefix so mcp-go still sees /mcp, /mcp/sse, etc.
			prefix := "/" + routeToken
			mux := http.NewServeMux()
			mux.Handle(prefix+"/", http.StripPrefix(prefix, httpHandler))
			handler = mux
		} else {
			handler = httpHandler
		}
		handler = httpLoggingMiddleware(handler, routeToken, multiTenant)
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
		srv := &http.Server{
			Addr:    addr,
			Handler: handler,
		}
		if tlsCert != "" && tlsKey != "" {
			return srv.ListenAndServeTLS(tlsCert, tlsKey)
		}
		return srv.ListenAndServe()
	default:
		return fmt.Errorf("unknown transport %q: must be stdio or http", transport)
	}
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

// TenantRoutingHandler extracts the Seerr API token from /{token}/mcp paths and
// injects it into the request context before forwarding to the MCP handler.
// Exported for testing.
func TenantRoutingHandler(mcpHandler http.Handler) http.Handler {
	return tenantRoutingHandler(mcpHandler)
}

func tenantRoutingHandler(mcpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Expect /{token}/mcp or /{token}/mcp/...
		path := strings.TrimPrefix(r.URL.Path, "/")
		slash := strings.Index(path, "/")
		if slash < 0 {
			http.NotFound(w, r)
			return
		}
		token, rest := path[:slash], path[slash:] // rest = "/mcp" or "/mcp/..."
		if token == "" || !strings.HasPrefix(rest, "/mcp") {
			http.NotFound(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), apiKeyCtxKey, token)
		r2 := r.Clone(ctx)
		r2.URL.Path = rest
		mcpHandler.ServeHTTP(w, r2)
	})
}
