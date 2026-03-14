package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	api "seer-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerBlocklistTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("blocklist_list",
			mcp.WithDescription("List all blocklisted media items"),
			mcp.WithNumber("take", mcp.Description("Number of results to return")),
			mcp.WithNumber("skip", mcp.Description("Number of results to skip")),
		),
		BlocklistListHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("blocklist_add",
			mcp.WithDescription("Add a media item to the blocklist"),
			mcp.WithNumber("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
			mcp.WithString("title", mcp.Description("Media title")),
		),
		BlocklistAddHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("blocklist_remove",
			mcp.WithDescription("Remove a media item from the blocklist"),
			mcp.WithString("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
		),
		BlocklistRemoveHandler(client, ctx),
	)
}

func BlocklistListHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.BlocklistAPI.BlocklistGet(ctx)
		if take := req.GetFloat("take", 0); take > 0 {
			r = r.Take(float32(take))
		}
		if skip := req.GetFloat("skip", 0); skip > 0 {
			r = r.Skip(float32(skip))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("BlocklistGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func BlocklistAddHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
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
		_, err = client.BlocklistAPI.BlocklistPost(ctx).Blocklist(body).Execute()
		if err != nil {
			return nil, fmt.Errorf("BlocklistPost failed: %w", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}

func BlocklistRemoveHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tmdbId, err := req.RequireString("tmdbId")
		if err != nil {
			return nil, err
		}
		_, err = client.BlocklistAPI.BlocklistTmdbIdDelete(ctx, tmdbId).Execute()
		if err != nil {
			return nil, fmt.Errorf("BlocklistTmdbIdDelete failed: %w", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}
