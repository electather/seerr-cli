package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPersonTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("person_get",
			mcp.WithDescription("Get person details by TMDB ID"),
			mcp.WithNumber("personId", mcp.Required(), mcp.Description("TMDB person ID")),
		),
		PersonGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("person_credits",
			mcp.WithDescription("Get combined credits for a person"),
			mcp.WithNumber("personId", mcp.Required(), mcp.Description("TMDB person ID")),
		),
		PersonCreditsHandler(),
	)
}

func PersonGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		personId, err := req.RequireFloat("personId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.PersonAPI.PersonPersonIdGet(callCtx, float32(personId)).Execute()
		if err != nil {
			return apiToolError("PersonPersonIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func PersonCreditsHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		personId, err := req.RequireFloat("personId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.PersonAPI.PersonPersonIdCombinedCreditsGet(callCtx, float32(personId)).Execute()
		if err != nil {
			return apiToolError("PersonPersonIdCombinedCreditsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
