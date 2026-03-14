package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerPersonTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("person_get",
			mcp.WithDescription("Get person details by TMDB ID"),
			mcp.WithNumber("personId", mcp.Required(), mcp.Description("TMDB person ID")),
		),
		PersonGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("person_credits",
			mcp.WithDescription("Get combined credits for a person"),
			mcp.WithNumber("personId", mcp.Required(), mcp.Description("TMDB person ID")),
		),
		PersonCreditsHandler(client, ctx),
	)
}

func PersonGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		personId, err := req.RequireFloat("personId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.PersonAPI.PersonPersonIdGet(ctx, float32(personId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("PersonPersonIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func PersonCreditsHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		personId, err := req.RequireFloat("personId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.PersonAPI.PersonPersonIdCombinedCreditsGet(ctx, float32(personId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("PersonPersonIdCombinedCreditsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
