package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerAuthTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("auth_me",
			mcp.WithDescription("Get the currently authenticated user's profile, permissions, and quota"),
			mcp.WithReadOnlyHintAnnotation(true),
		),
		AuthMeHandler(),
	)
}

// AuthMeHandler returns a tool handler for GET /auth/me.
func AuthMeHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.AuthAPI.AuthMeGet(callCtx).Execute()
		if err != nil {
			return apiToolError("AuthMeGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
