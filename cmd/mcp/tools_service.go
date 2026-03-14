package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerServiceTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("service_radarr_list",
			mcp.WithDescription("List configured Radarr instances"),
		),
		ServiceRadarrListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("service_sonarr_list",
			mcp.WithDescription("List configured Sonarr instances"),
		),
		ServiceSonarrListHandler(client, ctx),
	)
}

func ServiceRadarrListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.ServiceAPI.ServiceRadarrGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("ServiceRadarrGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func ServiceSonarrListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		res, _, err := client.ServiceAPI.ServiceSonarrGet(ctx).Execute()
		if err != nil {
			return nil, fmt.Errorf("ServiceSonarrGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
