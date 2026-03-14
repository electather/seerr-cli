package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	api "seer-cli/pkg/api"
)

func registerMoviesTools(s *server.MCPServer, client *api.APIClient, ctx context.Context) {
	s.AddTool(
		mcp.NewTool("movies_get",
			mcp.WithDescription("Get movie details by TMDB ID"),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
		),
		MoviesGetHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("movies_recommendations",
			mcp.WithDescription("Get movie recommendations for a given movie"),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		MoviesRecommendationsHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("movies_similar",
			mcp.WithDescription("Get similar movies for a given movie"),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		MoviesSimilarHandler(client, ctx),
	)

	s.AddTool(
		mcp.NewTool("movies_ratings",
			mcp.WithDescription("Get ratings for a given movie"),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
		),
		MoviesRatingsHandler(client, ctx),
	)
}

func MoviesGetHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.MoviesAPI.MovieMovieIdGet(ctx, float32(movieId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("MovieMovieIdGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesRecommendationsHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		r := client.MoviesAPI.MovieMovieIdRecommendationsGet(ctx, float32(movieId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("MovieMovieIdRecommendationsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesSimilarHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		r := client.MoviesAPI.MovieMovieIdSimilarGet(ctx, float32(movieId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return nil, fmt.Errorf("MovieMovieIdSimilarGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesRatingsHandler(client *api.APIClient, ctx context.Context) server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		res, _, err := client.MoviesAPI.MovieMovieIdRatingsGet(ctx, float32(movieId)).Execute()
		if err != nil {
			return nil, fmt.Errorf("MovieMovieIdRatingsGet failed: %w", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
