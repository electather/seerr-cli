package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerSearchTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("search_multi",
			mcp.WithDescription("Search for movies, TV shows, and people"),
			mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchMultiHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("search_discover_movies",
			mcp.WithDescription("Discover movies"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverMoviesHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("search_discover_tv",
			mcp.WithDescription("Discover TV shows"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverTVHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("search_trending",
			mcp.WithDescription("Get trending movies and TV shows"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchTrendingHandler(client, ctx),
	)
}

func SearchMultiHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return nil, fmt.Errorf("query is required")
		}
		r := client.SearchAPI.SearchGet(ctx).Query(query)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("SearchGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchDiscoverMoviesHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.SearchAPI.DiscoverMoviesGet(ctx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("DiscoverMoviesGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchDiscoverTVHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.SearchAPI.DiscoverTvGet(ctx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("DiscoverTvGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchTrendingHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		r := client.SearchAPI.DiscoverTrendingGet(ctx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("DiscoverTrendingGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
