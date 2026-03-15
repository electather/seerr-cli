package mcp

import (
	"context"
	"encoding/json"

	api "seer-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerBlocklistTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("blocklist_list",
			mcp.WithDescription("List all blocklisted media items"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		BlocklistListHandler(),
	)

	s.AddTool(
		mcp.NewTool("blocklist_add",
			mcp.WithDescription("Add a media item to the blocklist"),
			mcp.WithNumber("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
			mcp.WithString("title", mcp.Description("Media title")),
		),
		BlocklistAddHandler(),
	)

	s.AddTool(
		mcp.NewTool("blocklist_remove",
			mcp.WithDescription("Remove a media item from the blocklist"),
			mcp.WithString("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
		),
		BlocklistRemoveHandler(),
	)
}

func BlocklistListHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.BlocklistAPI.BlocklistGet(callCtx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("BlocklistGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func BlocklistAddHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tmdbId, err := req.RequireFloat("tmdbId")
		if err != nil {
			return nil, err
		}
		tmdbIdFloat := float32(tmdbId)
		body := api.Blocklist{
			TmdbId: &tmdbIdFloat,
		}
		if title := req.GetString("title", ""); title != "" {
			body.Title = &title
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		_, err = client.BlocklistAPI.BlocklistPost(callCtx).Blocklist(body).Execute()
		if err != nil {
			return apiToolError("BlocklistPost failed", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}

func BlocklistRemoveHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tmdbId, err := req.RequireString("tmdbId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		_, err = client.BlocklistAPI.BlocklistTmdbIdDelete(callCtx, tmdbId).Execute()
		if err != nil {
			return apiToolError("BlocklistTmdbIdDelete failed", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}
