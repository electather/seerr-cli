package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerStatusTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("status_system",
			mcp.WithDescription("Get the system status of the Seer API"),
		),
		StatusSystemHandler(client, ctx),
	)
}

func StatusSystemHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.PublicAPI.StatusGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("StatusGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
