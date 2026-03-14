package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerTVTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("tv_get",
			mcp.WithDescription("Get TV show details by TMDB ID"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
		),
		TVGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("tv_season",
			mcp.WithDescription("Get season details for a TV show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("seasonNumber", mcp.Required(), mcp.Description("Season number")),
		),
		TVSeasonHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("tv_recommendations",
			mcp.WithDescription("Get TV show recommendations for a given show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		TVRecommendationsHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("tv_similar",
			mcp.WithDescription("Get similar TV shows for a given show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		TVSimilarHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("tv_ratings",
			mcp.WithDescription("Get ratings for a given TV show"),
			mcp.WithNumber("tvId", mcp.Required(), mcp.Description("TMDB TV show ID")),
		),
		TVRatingsHandler(client, ctx),
	)
}

func TVGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.TvAPI.TvTvIdGet(ctx, float32(tvId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("TvTvIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVSeasonHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		seasonNumber, err := req.RequireFloat("seasonNumber")
		if err != nil {
			return nil, err
		}
		res, _, err := client.TvAPI.TvTvIdSeasonSeasonNumberGet(ctx, float32(tvId), float32(seasonNumber)).Execute()
		if err != nil {
			return nil, fmt.Errorf("TvTvIdSeasonSeasonNumberGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVRecommendationsHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		r := client.TvAPI.TvTvIdRecommendationsGet(ctx, float32(tvId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("TvTvIdRecommendationsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVSimilarHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		r := client.TvAPI.TvTvIdSimilarGet(ctx, float32(tvId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("TvTvIdSimilarGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func TVRatingsHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		tvId, err := req.RequireFloat("tvId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.TvAPI.TvTvIdRatingsGet(ctx, float32(tvId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("TvTvIdRatingsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
