package seerrclient

import (
	"context"
	"net/url"
	"strconv"
)

// SearchMulti searches for movies, TV shows, and people. It uses a raw HTTP
// request to work around the generated client's broken union-type unmarshal
// where TV results are incorrectly deserialised as PersonResult.
func (c *Client) SearchMulti(query string, page int, language string) ([]byte, error) {
	params := url.Values{}
	params.Set("query", query)
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if language != "" {
		params.Set("language", language)
	}
	return c.RawGet("/search", params)
}

// SearchMultiCtx is like SearchMulti but accepts an explicit context.
func (c *Client) SearchMultiCtx(ctx context.Context, query string, page int, language string) ([]byte, error) {
	params := url.Values{}
	params.Set("query", query)
	if page > 0 {
		params.Set("page", strconv.Itoa(page))
	}
	if language != "" {
		params.Set("language", language)
	}
	return c.RawGetCtx(ctx, "/search", params)
}

// DiscoverMovies runs the /discover/movies endpoint with arbitrary filter
// parameters. All numeric values should be pre-formatted as strings in params.
func (c *Client) DiscoverMovies(params url.Values) ([]byte, error) {
	return c.RawGet("/discover/movies", params)
}

// DiscoverTV runs the /discover/tv endpoint with arbitrary filter parameters.
func (c *Client) DiscoverTV(params url.Values) ([]byte, error) {
	return c.RawGet("/discover/tv", params)
}

// DiscoverTrending runs the /discover/trending endpoint. Default values for
// mediaType ("all") and timeWindow ("day") are stripped from params before the
// request is sent because some Seerr server versions reject them with HTTP 400.
func (c *Client) DiscoverTrending(params url.Values) ([]byte, error) {
	// Clone params so we don't mutate the caller's map.
	p := url.Values{}
	for k, vs := range params {
		p[k] = vs
	}
	if p.Get("mediaType") == "all" {
		p.Del("mediaType")
	}
	if p.Get("timeWindow") == "day" {
		p.Del("timeWindow")
	}
	return c.RawGet("/discover/trending", p)
}

// DiscoverTrendingCtx is like DiscoverTrending but accepts an explicit context.
func (c *Client) DiscoverTrendingCtx(ctx context.Context, params url.Values) ([]byte, error) {
	p := url.Values{}
	for k, vs := range params {
		p[k] = vs
	}
	if p.Get("mediaType") == "all" {
		p.Del("mediaType")
	}
	if p.Get("timeWindow") == "day" {
		p.Del("timeWindow")
	}
	return c.RawGetCtx(ctx, "/discover/trending", p)
}
