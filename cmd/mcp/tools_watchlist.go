package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerWatchlistTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("watchlist_add",
			mcp.WithDescription("Add a media item to the watchlist"),
			mcp.WithNumber("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
			mcp.WithString("title", mcp.Required(), mcp.Description("Media title")),
			mcp.WithString("mediaType", mcp.Required(), mcp.Description("Media type: movie or tv")),
		),
		WatchlistAddHandler(),
	)

	s.AddTool(
		mcp.NewTool("watchlist_remove",
			mcp.WithDescription("Remove a media item from the watchlist"),
			mcp.WithString("tmdbId", mcp.Required(), mcp.Description("TMDB media ID")),
		),
		WatchlistRemoveHandler(),
	)
}

func WatchlistAddHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tmdbId, err := req.RequireFloat("tmdbId")
		if err != nil {
			return nil, err
		}
		title, err := req.RequireString("title")
		if err != nil {
			return nil, err
		}
		mediaType, err := req.RequireString("mediaType")
		if err != nil {
			return nil, err
		}
		tmdbIdFloat := float32(tmdbId)
		body := api.Watchlist{
			TmdbId: &tmdbIdFloat,
			Title:  &title,
			Type:   &mediaType,
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.WatchlistAPI.WatchlistPost(callCtx).Watchlist(body).Execute()
		if err != nil {
			return apiToolError("WatchlistPost failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func WatchlistRemoveHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tmdbId, err := req.RequireString("tmdbId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		_, err = client.WatchlistAPI.WatchlistTmdbIdDelete(callCtx, tmdbId).Execute()
		if err != nil {
			return apiToolError("WatchlistTmdbIdDelete failed", err)
		}
		return mcp.NewToolResultText(`{"status":"ok"}`), nil
	}
}
