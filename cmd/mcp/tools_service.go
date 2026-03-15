package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerServiceTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("service_radarr_list",
			mcp.WithDescription("List configured Radarr instances"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		ServiceRadarrListHandler(),
	)

	s.AddTool(
		mcp.NewTool("service_sonarr_list",
			mcp.WithDescription("List configured Sonarr instances"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		ServiceSonarrListHandler(),
	)
}

func ServiceRadarrListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.ServiceAPI.ServiceRadarrGet(callCtx).Execute()
		if err != nil {
			return apiToolError("ServiceRadarrGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func ServiceSonarrListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.ServiceAPI.ServiceSonarrGet(callCtx).Execute()
		if err != nil {
			return apiToolError("ServiceSonarrGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
