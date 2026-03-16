package mcp

import (
	"context"
	"encoding/json"
	"strings"
	"sync"

	api "seerr-cli/pkg/api"
)

// tmdbImageBase is the base URL for TMDB poster and backdrop images at w500 size.
const tmdbImageBase = "https://image.tmdb.org/t/p/w500"

// GenreMap maps a TMDB genre ID to its human-readable name.
type GenreMap map[int]string

// FetchMovieGenres fetches the TMDB movie genre list and returns a lookup map.
// Returns nil on error; genre enrichment is best-effort.
func FetchMovieGenres(ctx context.Context, client *api.APIClient) GenreMap {
	genres, _, err := client.TmdbAPI.GenresMovieGet(ctx).Execute()
	if err != nil {
		return nil
	}
	m := make(GenreMap, len(genres))
	for _, g := range genres {
		if g.Id != nil && g.Name != nil {
			m[int(*g.Id)] = *g.Name
		}
	}
	return m
}

// FetchTVGenres fetches the TMDB TV genre list and returns a lookup map.
// Returns nil on error; genre enrichment is best-effort.
func FetchTVGenres(ctx context.Context, client *api.APIClient) GenreMap {
	genres, _, err := client.TmdbAPI.GenresTvGet(ctx).Execute()
	if err != nil {
		return nil
	}
	m := make(GenreMap, len(genres))
	for _, g := range genres {
		if g.Id != nil && g.Name != nil {
			m[int(*g.Id)] = *g.Name
		}
	}
	return m
}

// MergeGenreMaps merges multiple GenreMaps into one.
// TMDB genre IDs are consistent across media types, so conflicts are rare.
func MergeGenreMaps(maps ...GenreMap) GenreMap {
	merged := make(GenreMap)
	for _, m := range maps {
		for id, name := range m {
			merged[id] = name
		}
	}
	return merged
}

// FetchAllGenres fetches movie and TV genre maps concurrently and returns their union.
func FetchAllGenres(ctx context.Context, client *api.APIClient) GenreMap {
	var (
		movieGenres, tvGenres GenreMap
		wg                    sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		movieGenres = FetchMovieGenres(ctx, client)
	}()
	go func() {
		defer wg.Done()
		tvGenres = FetchTVGenres(ctx, client)
	}()
	wg.Wait()
	return MergeGenreMaps(movieGenres, tvGenres)
}

// EnrichMediaMap enriches a single JSON map representing a movie or TV result.
// It adds a genreNames field resolved from genres and expands image paths to
// full TMDB URLs. The map is modified in-place.
func EnrichMediaMap(m map[string]interface{}, genres GenreMap) {
	if len(genres) > 0 {
		if ids, ok := m["genreIds"].([]interface{}); ok && len(ids) > 0 {
			names := make([]string, 0, len(ids))
			for _, v := range ids {
				if id, ok := v.(float64); ok {
					if name := genres[int(id)]; name != "" {
						names = append(names, name)
					}
				}
			}
			if len(names) > 0 {
				m["genreNames"] = names
			}
		}
	}
	for _, field := range []string{"posterPath", "backdropPath"} {
		if path, ok := m[field].(string); ok && path != "" && !strings.HasPrefix(path, "http") {
			m[field] = tmdbImageBase + path
		}
	}
}

// EnrichResultsPage enriches every item in the results array of a paged API
// response. The input bytes are returned unchanged on any parse error.
func EnrichResultsPage(data []byte, genres GenreMap) ([]byte, error) {
	var page map[string]interface{}
	if err := json.Unmarshal(data, &page); err != nil {
		return data, nil
	}
	results, ok := page["results"].([]interface{})
	if !ok {
		return data, nil
	}
	for _, item := range results {
		if m, ok := item.(map[string]interface{}); ok {
			EnrichMediaMap(m, genres)
		}
	}
	return json.Marshal(page)
}
