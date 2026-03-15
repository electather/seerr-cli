package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerMoviesTools(s *server.MCPServer) {
	s.AddTool(
		mcp.NewTool("movies_get",
			mcp.WithDescription("Get movie details by TMDB ID"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
		),
		MoviesGetHandler(),
	)

	s.AddTool(
		mcp.NewTool("movies_recommendations",
			mcp.WithDescription("Get movie recommendations for a given movie"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		MoviesRecommendationsHandler(),
	)

	s.AddTool(
		mcp.NewTool("movies_similar",
			mcp.WithDescription("Get similar movies for a given movie"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
			mcp.WithNumber("page", mcp.Description("Page number")),
		),
		MoviesSimilarHandler(),
	)

	s.AddTool(
		mcp.NewTool("movies_ratings",
			mcp.WithDescription("Get ratings for a given movie"),
			mcp.WithDestructiveHintAnnotation(false),
			mcp.WithReadOnlyHintAnnotation(true),
			mcp.WithIdempotentHintAnnotation(true),
			mcp.WithNumber("movieId", mcp.Required(), mcp.Description("TMDB movie ID")),
		),
		MoviesRatingsHandler(),
	)
}

func MoviesGetHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.MoviesAPI.MovieMovieIdGet(callCtx, float32(movieId)).Execute()
		if err != nil {
			return apiToolError("MovieMovieIdGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesRecommendationsHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.MoviesAPI.MovieMovieIdRecommendationsGet(callCtx, float32(movieId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("MovieMovieIdRecommendationsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesSimilarHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		r := client.MoviesAPI.MovieMovieIdSimilarGet(callCtx, float32(movieId))
		if page := req.GetFloat("page", 0); page > 0 {
			r = r.Page(float32(page))
		}
		res, _, err := r.Execute()
		if err != nil {
			return apiToolError("MovieMovieIdSimilarGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}

func MoviesRatingsHandler() server.ToolHandlerFunc {
	return func(callCtx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		movieId, err := req.RequireFloat("movieId")
		if err != nil {
			return nil, err
		}
		client := newAPIClientWithKey(apiKeyFromContext(callCtx))
		res, _, err := client.MoviesAPI.MovieMovieIdRatingsGet(callCtx, float32(movieId)).Execute()
		if err != nil {
			return apiToolError("MovieMovieIdRatingsGet failed", err)
		}
		b, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		return mcp.NewToolResultText(string(b)), nil
	}
}
