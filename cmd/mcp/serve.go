package mcp

import (
	"context"
	"crypto/subtle"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the MCP server",
	Long:  `Start a Model Context Protocol server that exposes the Seer API as tools.`,
	Example: `  # Start with stdio transport (default, for Claude Desktop)
  seer-cli mcp serve

  # Start over HTTP with Bearer token auth
  seer-cli mcp serve --transport http --auth-token mysecret

  # Start over HTTPS with TLS
  seer-cli mcp serve --transport http --auth-token mysecret --tls-cert /path/to/cert.pem --tls-key /path/to/key.pem

  # Start over HTTP with a secret path prefix (for clients that cannot send custom headers)
  seer-cli mcp serve --transport http --route-token abc123 --no-auth
  # MCP endpoint becomes: http://localhost:8811/abc123/mcp

  # Enable CORS for browser-based clients (e.g. claude.ai)
  seer-cli mcp serve --transport http --route-token abc123 --no-auth --cors

  # Combine both auth methods for defense in depth
  seer-cli mcp serve --transport http --route-token abc123 --auth-token mysecret

  # Start over HTTP without auth (insecure, not recommended)
  seer-cli mcp serve --transport http --no-auth

  # Start in multi-tenant mode (per-user API keys in URL path)
  seer-cli mcp serve --transport http --no-auth --multi-tenant`,
	RunE: runServe,
}

func init() {
	serveCmd.Flags().String("transport", "stdio", "Transport protocol: stdio or http (env: SEER_MCP_TRANSPORT)")
	serveCmd.Flags().String("addr", ":8811", "HTTP bind address (http transport only) (env: SEER_MCP_ADDR)")
	serveCmd.Flags().String("auth-token", "", "Bearer token required for HTTP transport (env: SEER_MCP_AUTH_TOKEN)")
	serveCmd.Flags().Bool("no-auth", false, "Disable authentication (insecure — must be explicit) (env: SEER_MCP_NO_AUTH)")
	serveCmd.Flags().String("route-token", "", "Secret path prefix for the MCP endpoint (e.g. 'abc123' → /abc123/mcp). Useful for clients that cannot send custom headers (env: SEER_MCP_ROUTE_TOKEN)")
	serveCmd.Flags().String("tls-cert", "", "Path to TLS certificate file (env: SEER_MCP_TLS_CERT)")
	serveCmd.Flags().String("tls-key", "", "Path to TLS private key file (env: SEER_MCP_TLS_KEY)")
	serveCmd.Flags().Bool("cors", false, "Enable CORS headers (required for browser-based clients such as claude.ai) (env: SEER_MCP_CORS)")
	serveCmd.Flags().Bool("multi-tenant", false, "Route /{seer-api-token}/mcp for per-user API keys (HTTP transport only)")
	viper.BindPFlag("mcp_transport", serveCmd.Flags().Lookup("transport"))
	viper.BindPFlag("mcp_addr", serveCmd.Flags().Lookup("addr"))
	viper.BindPFlag("mcp_auth_token", serveCmd.Flags().Lookup("auth-token"))
	viper.BindPFlag("mcp_no_auth", serveCmd.Flags().Lookup("no-auth"))
	viper.BindPFlag("mcp_route_token", serveCmd.Flags().Lookup("route-token"))
	viper.BindPFlag("mcp_tls_cert", serveCmd.Flags().Lookup("tls-cert"))
	viper.BindPFlag("mcp_tls_key", serveCmd.Flags().Lookup("tls-key"))
	viper.BindPFlag("mcp_cors", serveCmd.Flags().Lookup("cors"))
	viper.BindPFlag("mcp_multi_tenant", serveCmd.Flags().Lookup("multi-tenant"))
	Cmd.AddCommand(serveCmd)
}

func runServe(_ *cobra.Command, args []string) error {
	transport := viper.GetString("mcp_transport")
	addr := viper.GetString("mcp_addr")
	authToken := viper.GetString("mcp_auth_token")
	routeToken := viper.GetString("mcp_route_token")
	noAuth := viper.GetBool("mcp_no_auth")
	tlsCert := viper.GetString("mcp_tls_cert")
	tlsKey := viper.GetString("mcp_tls_key")
	cors := viper.GetBool("mcp_cors")
	multiTenant := viper.GetBool("mcp_multi_tenant")

	if transport == "http" && authToken == "" && routeToken == "" && !noAuth {
		return fmt.Errorf("HTTP transport requires --auth-token, --route-token, or --no-auth (insecure) to be set explicitly")
	}

	if multiTenant && transport != "http" {
		return fmt.Errorf("--multi-tenant requires --transport http")
	}

	verbose := viper.GetBool("verbose")

	s := server.NewMCPServer("seer-mcp", "1.0.0")

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

	switch transport {
	case "stdio":
		if verbose {
			seerServer := viper.GetString("server")
			fmt.Fprintf(os.Stderr, "seer-mcp: transport=stdio\n")
			fmt.Fprintf(os.Stderr, "seer-mcp: seer API → %s\n", seerServer)
			fmt.Fprintf(os.Stderr, "seer-mcp: tools registered: 44\n")
			fmt.Fprintf(os.Stderr, "seer-mcp: waiting for MCP client on stdin…\n")
		} else {
			fmt.Fprintf(os.Stderr, "seer-mcp: ready (stdio) — waiting for MCP client on stdin\n")
		}
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
			mcpPath = "/{seer-api-token}/mcp"
		} else if routeToken != "" {
			mcpPath = "/" + routeToken + "/mcp"
		}
		endpoint := fmt.Sprintf("%s://%s%s", scheme, host, mcpPath)

		fmt.Fprintf(os.Stderr, "seer-mcp: listening on %s\n", endpoint)
		fmt.Fprintf(os.Stderr, "\nConfigure your MCP client:\n")
		fmt.Fprintf(os.Stderr, "  URL:  %s\n", endpoint)
		if authToken != "" {
			fmt.Fprintf(os.Stderr, "  Authorization: Bearer %s\n", authToken)
		} else {
			fmt.Fprintf(os.Stderr, "  Authorization: none\n")
		}

		if verbose {
			seerServer := viper.GetString("server")
			fmt.Fprintf(os.Stderr, "\nVerbose:\n")
			fmt.Fprintf(os.Stderr, "  Seer API → %s\n", seerServer)
			fmt.Fprintf(os.Stderr, "  Bind addr: %s\n", addr)
			fmt.Fprintf(os.Stderr, "  TLS: %v\n", tlsCert != "")
			fmt.Fprintf(os.Stderr, "  Auth token: %v\n", authToken != "")
			fmt.Fprintf(os.Stderr, "  Route token: %v\n", routeToken != "")
			fmt.Fprintf(os.Stderr, "  CORS: %v\n", cors)
			fmt.Fprintf(os.Stderr, "  Tools registered: 44\n")
			fmt.Fprintf(os.Stderr, "  Multi-tenant: %v\n", multiTenant)
		}

		fmt.Fprintf(os.Stderr, "\n")

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
		if authToken != "" {
			handler = bearerAuthMiddleware(authToken, handler)
		}
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

// TenantRoutingHandler extracts the Seer API token from /{token}/mcp paths and
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
