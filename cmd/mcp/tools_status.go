package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerStatusTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("status_system",
			mcp.WithDescription("Get the system status of the Seer API"),
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
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
