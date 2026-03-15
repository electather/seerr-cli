package mcp

import (
	"context"
	api "seerr-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerStatusTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("status_system",
			mcp.WithDescription("Get the system status of the Seerr API"),
			mcp.WithOpenWorldHintAnnotation(true),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithOutputSchema[api.StatusGet200Response](),
		),
		StatusSystemHandler(),
	)
}

func StatusSystemHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.PublicAPI.StatusGet(callCtx).Execute()
		if err != nil {
			return apiToolError("StatusGet failed", err)
		}
		return mcp.NewToolResultJSON(res)
	}
}
