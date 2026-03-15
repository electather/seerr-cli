package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTVTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("tv_get",
			mcp.WithDescription("Get TV show details by TMDB ID"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
		),
		TVGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("tv_season",
			mcp.WithDescription("Get season details for a TV show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("seasonNumber", mcp.Required(), mcp.Description("Season number")),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		TVSeasonHandler(),
	)

	s.AddTool(
		mcp.NewTool("tv_recommendations",
			mcp.WithDescription("Get TV show recommendations for a given show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		TVRecommendationsHandler(),
	)

	s.AddTool(
		mcp.NewTool("tv_similar",
			mcp.WithDescription("Get similar TV shows for a given show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		TVSimilarHandler(),
	)

	s.AddTool(
		mcp.NewTool("tv_ratings",
			mcp.WithDescription("Get ratings for a given TV show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
		),
		TVRatingsHandler(),
	)
}

func TVGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.TvAPI.TvTvIdGet(callCtx, float32(tvId)).Execute()
		if err != nil {
			return apiToolError("TvTvIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVSeasonHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		seasonNumber, err := req.RequireFloat("seasonNumber")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.TvAPI.TvTvIdSeasonSeasonNumberGet(callCtx, float32(tvId), float32(seasonNumber)).Execute()
		if err != nil {
			return apiToolError("TvTvIdSeasonSeasonNumberGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVRecommendationsHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.TvAPI.TvTvIdRecommendationsGet(callCtx, float32(tvId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("TvTvIdRecommendationsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVSimilarHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.TvAPI.TvTvIdSimilarGet(callCtx, float32(tvId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("TvTvIdSimilarGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVRatingsHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.TvAPI.TvTvIdRatingsGet(callCtx, float32(tvId)).Execute()
		if err != nil {
			return apiToolError("TvTvIdRatingsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
