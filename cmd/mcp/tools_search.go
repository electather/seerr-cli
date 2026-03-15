package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerSearchTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("search_multi",
			mcp.WithDescription("Search for movies, TV shows, and people"),
			mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchMultiHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_discover_movies",
			mcp.WithDescription("Discover movies"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverMoviesHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_discover_tv",
			mcp.WithDescription("Discover TV shows"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchDiscoverTVHandler(),
	)

	s.AddTool(
		mcp.NewTool("search_trending",
			mcp.WithDescription("Get trending movies and TV shows"),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		SearchTrendingHandler(),
	)
}

func SearchMultiHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query := req.GetString("query", "")
		if query == "" {
			return nil, fmt.Errorf("query is required")
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.SearchGet(callCtx).Query(query)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("SearchGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchDiscoverMoviesHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.DiscoverMoviesGet(callCtx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("DiscoverMoviesGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchDiscoverTVHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.DiscoverTvGet(callCtx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("DiscoverTvGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func SearchTrendingHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.SearchAPI.DiscoverTrendingGet(callCtx)
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("DiscoverTrendingGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
