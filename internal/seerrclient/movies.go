package seerrclient

import (
	"context"
	"net/http"

	api "seerr-cli/pkg/api"
)

// MovieGet returns details for a single movie. Pass an empty language string
// to use the server default.
func (c *Client) MovieGet(id int, language string) (*api.MovieDetails, *http.Response, error) {
	req := c.api.MoviesAPI.MovieMovieIdGet(c.ctx, float32(id))
	if language != "" {
		req = req.Language(language)
	}
	return req.Execute()
}

// MovieGetCtx is like MovieGet but accepts an explicit context.
func (c *Client) MovieGetCtx(ctx context.Context, id int, language string) (*api.MovieDetails, *http.Response, error) {
	req := c.api.MoviesAPI.MovieMovieIdGet(ctx, float32(id))
	if language != "" {
		req = req.Language(language)
	}
	return req.Execute()
}

// MovieRecommendations returns recommended movies for the given movie ID.
// Pass page <= 0 to use the server default pagination.
func (c *Client) MovieRecommendations(id int, page int, language string) (*api.DiscoverMoviesGet200Response, *http.Response, error) {
	req := c.api.MoviesAPI.MovieMovieIdRecommendationsGet(c.ctx, float32(id))
	if page > 0 {
		req = req.Page(float32(page))
	}
	if language != "" {
		req = req.Language(language)
	}
	return req.Execute()
}

// MovieSimilar returns movies similar to the given movie ID.
func (c *Client) MovieSimilar(id int, page int, language string) (*api.DiscoverMoviesGet200Response, *http.Response, error) {
	req := c.api.MoviesAPI.MovieMovieIdSimilarGet(c.ctx, float32(id))
	if page > 0 {
		req = req.Page(float32(page))
	}
	if language != "" {
		req = req.Language(language)
	}
	return req.Execute()
}

// MovieRatings returns ratings data for the given movie ID.
func (c *Client) MovieRatings(id int) (*api.MovieMovieIdRatingsGet200Response, *http.Response, error) {
	return c.api.MoviesAPI.MovieMovieIdRatingsGet(c.ctx, float32(id)).Execute()
}

// MovieRatingsCombined returns combined RT and IMDB ratings for the given movie ID.
func (c *Client) MovieRatingsCombined(id int) (*api.MovieMovieIdRatingscombinedGet200Response, *http.Response, error) {
	return c.api.MoviesAPI.MovieMovieIdRatingscombinedGet(c.ctx, float32(id)).Execute()
}
