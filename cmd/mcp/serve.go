package mcp

import (
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

  # Start over HTTP without auth (insecure, not recommended)
  seer-cli mcp serve --transport http --no-auth`,
	RunE: runServe,
}

func init() {
	serveCmd.Flags().String("transport", "stdio", "Transport protocol: stdio or http")
	serveCmd.Flags().String("addr", ":8811", "HTTP bind address (http transport only)")
	serveCmd.Flags().String("auth-token", "", "Bearer token required for HTTP transport")
	serveCmd.Flags().Bool("no-auth", false, "Disable authentication (insecure — must be explicit)")
	serveCmd.Flags().String("tls-cert", "", "Path to TLS certificate file")
	serveCmd.Flags().String("tls-key", "", "Path to TLS private key file")
	Cmd.AddCommand(serveCmd)
}

func runServe(cmd *cobra.Command, args []string) error {
	transport, _ := cmd.Flags().GetString("transport")
	addr, _ := cmd.Flags().GetString("addr")
	authToken, _ := cmd.Flags().GetString("auth-token")
	noAuth, _ := cmd.Flags().GetBool("no-auth")
	tlsCert, _ := cmd.Flags().GetString("tls-cert")
	tlsKey, _ := cmd.Flags().GetString("tls-key")

	if transport == "http" && authToken == "" && !noAuth {
		return fmt.Errorf("HTTP transport requires --auth-token or --no-auth (insecure) to be set explicitly")
	}

	verbose := viper.GetBool("verbose")
	client, ctx := newAPIClient()

	s := server.NewMCPServer("seer-mcp", "1.0.0")

	registerStatusTools(s, client, ctx)
	registerSearchTools(s, client, ctx)
	registerMoviesTools(s, client, ctx)
	registerTVTools(s, client, ctx)
	registerRequestTools(s, client, ctx)
	registerMediaTools(s, client, ctx)
	registerUsersTools(s, client, ctx)
	registerIssueTools(s, client, ctx)
	registerPersonTools(s, client, ctx)
	registerCollectionTools(s, client, ctx)
	registerServiceTools(s, client, ctx)
	registerSettingsTools(s, client, ctx)
	registerWatchlistTools(s, client, ctx)
	registerBlocklistTools(s, client, ctx)

	switch transport {
	case "stdio":
		if verbose {
			seerServer := viper.GetString("server")
			fmt.Fprintf(os.Stderr, "seer-mcp: transport=stdio\n")
			fmt.Fprintf(os.Stderr, "seer-mcp: seer API → %s\n", seerServer)
			fmt.Fprintf(os.Stderr, "seer-mcp: tools registered: 43\n")
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
		endpoint := fmt.Sprintf("%s://%s/mcp", scheme, host)

		fmt.Fprintf(os.Stderr, "seer-mcp: listening on %s\n", endpoint)
		fmt.Fprintf(os.Stderr, "\nConfigure your MCP client:\n")
		fmt.Fprintf(os.Stderr, "  URL:  %s\n", endpoint)
		if authToken != "" {
			fmt.Fprintf(os.Stderr, "  Authorization: Bearer %s\n", authToken)
		} else {
			fmt.Fprintf(os.Stderr, "  Authorization: none (--no-auth)\n")
		}

		if verbose {
			seerServer := viper.GetString("server")
			fmt.Fprintf(os.Stderr, "\nVerbose:\n")
			fmt.Fprintf(os.Stderr, "  Seer API → %s\n", seerServer)
			fmt.Fprintf(os.Stderr, "  Bind addr: %s\n", addr)
			fmt.Fprintf(os.Stderr, "  TLS: %v\n", tlsCert != "")
			fmt.Fprintf(os.Stderr, "  Auth: %v\n", authToken != "")
			fmt.Fprintf(os.Stderr, "  Tools registered: 43\n")
		}

		fmt.Fprintf(os.Stderr, "\n")

		httpHandler := server.NewStreamableHTTPServer(s)
		var handler http.Handler = httpHandler
		if authToken != "" {
			handler = bearerAuthMiddleware(authToken, httpHandler)
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
