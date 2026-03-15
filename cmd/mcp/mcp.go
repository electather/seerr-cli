package mcp

import (
	"context"
	"fmt"
	"strings"

	api "seerr-cli/pkg/api"

	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

type contextKey string

const apiKeyCtxKey contextKey = "seerApiKey"

// APIKeyContextKey is exported for tests to inject a key directly.
const APIKeyContextKey = apiKeyCtxKey

func apiKeyFromContext(ctx context.Context) string {
	v, _ := ctx.Value(apiKeyCtxKey).(string)
	return v
}

var Cmd = &cobra.Command{
	Use:   "mcp",
	Short: "Run a Model Context Protocol (MCP) server exposing the Seerr API",
	Long:  `Start an MCP server that exposes the Seerr API as tools for use by AI agents.`,
	Example: `  # Start the MCP server using stdio transport (for Claude Desktop)
  seerr-cli mcp serve

  # Start the MCP server over HTTP with a Bearer token
  seerr-cli mcp serve --transport http --auth-token mysecret`,
}

// newAPIClientWithKey builds a client using apiKey, falling back to Viper when empty.
func newAPIClientWithKey(apiKey string) *api.APIClient {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("seer.server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	key := apiKey
	if key == "" {
		key = viper.GetString("seer.api_key")
	}
	if key != "" {
		configuration.AddDefaultHeader("X-Api-Key", key)
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}
	return api.NewAPIClient(configuration)
}

func newAPIClient() (*api.APIClient, context.Context) {
	return newAPIClientWithKey(""), context.Background()
}

// NewAPIClientForTest is an exported alias used by tests.
func NewAPIClientForTest() (*api.APIClient, context.Context) {
	return newAPIClient()
}

// apiToolError converts a Seerr API error into a visible MCP tool result error.
// Using mcp.NewToolResultError (IsError=true) instead of returning a Go error
// makes the actual Seerr error message visible in the MCP client (e.g. Claude)
// rather than the generic "Error occurred during tool execution" wrapper.
func apiToolError(label string, err error) (*mcplib.CallToolResult, error) {
	msg := fmt.Sprintf("%s: %v", label, err)
	if e, ok := err.(*api.GenericOpenAPIError); ok && len(e.Body()) > 0 {
		msg = fmt.Sprintf("%s: %s (HTTP %v)", label, e.Body(), err)
	}
	if mcpLog != nil {
		mcpLog.Warn("tool error", "label", label, "error", err)
	}
	return mcplib.NewToolResultError(msg), nil
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
