package mcp

import (
	"context"
	"encoding/json"
	"sync"

	api "seerr-cli/pkg/api"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerResources(s *server.MCPServer) {
	s.AddResource(
		mcp.NewResource("seerr://genres/movies", "Movie Genres",
			mcp.WithResourceDescription("TMDB movie genre ID to name map"),
			mcp.WithMIMEType("application/json"),
		),
		GenresMovieResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://genres/tv", "TV Genres",
			mcp.WithResourceDescription("TMDB TV genre ID to name map"),
			mcp.WithMIMEType("application/json"),
		),
		GenresTVResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://languages", "Languages",
			mcp.WithResourceDescription("ISO language codes and English names supported by TMDB"),
			mcp.WithMIMEType("application/json"),
		),
		LanguagesResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://regions", "Regions",
			mcp.WithResourceDescription("ISO region codes and English names supported by TMDB"),
			mcp.WithMIMEType("application/json"),
		),
		RegionsResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://certifications/movies", "Movie Certifications",
			mcp.WithResourceDescription("Content ratings by country for movies (G, PG, R, etc.)"),
			mcp.WithMIMEType("application/json"),
		),
		CertificationsMovieResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://certifications/tv", "TV Certifications",
			mcp.WithResourceDescription("Content ratings by country for TV shows"),
			mcp.WithMIMEType("application/json"),
		),
		CertificationsTVResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://services/radarr", "Radarr Services",
			mcp.WithResourceDescription("Configured Radarr instances with quality profiles and root folders"),
			mcp.WithMIMEType("application/json"),
		),
		RadarrServicesResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://services/sonarr", "Sonarr Services",
			mcp.WithResourceDescription("Configured Sonarr instances with quality profiles and root folders"),
			mcp.WithMIMEType("application/json"),
		),
		SonarrServicesResourceHandler(),
	)

	s.AddResource(
		mcp.NewResource("seerr://system/about", "System Info",
			mcp.WithResourceDescription("Seerr version, commit hash, and update availability"),
			mcp.WithMIMEType("application/json"),
		),
		SystemAboutResourceHandler(),
	)
}

// textResource wraps JSON bytes as a single TextResourceContents slice.
func textResource(uri string, data []byte) []mcp.ResourceContents {
	return []mcp.ResourceContents{mcp.TextResourceContents{
		URI:      uri,
		MIMEType: "application/json",
		Text:     string(data),
	}}
}

// GenresMovieResourceHandler returns a resource handler for the TMDB movie genre list.
func GenresMovieResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.TmdbAPI.GenresMovieGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// GenresTVResourceHandler returns a resource handler for the TMDB TV genre list.
func GenresTVResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.TmdbAPI.GenresTvGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// LanguagesResourceHandler returns a resource handler for the TMDB language list.
func LanguagesResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.TmdbAPI.LanguagesGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// RegionsResourceHandler returns a resource handler for the TMDB region list.
func RegionsResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.TmdbAPI.RegionsGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// CertificationsMovieResourceHandler returns a resource handler for movie content ratings.
func CertificationsMovieResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.OtherAPI.CertificationsMovieGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// CertificationsTVResourceHandler returns a resource handler for TV content ratings.
func CertificationsTVResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.OtherAPI.CertificationsTvGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// WatchProvidersMoviesResourceHandler returns a resource handler for movie streaming providers.
func WatchProvidersMoviesResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.OtherAPI.WatchprovidersMoviesGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// WatchProvidersTVResourceHandler returns a resource handler for TV streaming providers.
func WatchProvidersTVResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.OtherAPI.WatchprovidersTvGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// RadarrServicesResourceHandler returns a resource handler that lists all configured
// Radarr instances enriched with per-instance quality profiles and root folders.
func RadarrServicesResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		instances, _, err := client.ServiceAPI.ServiceRadarrGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		enriched := enrichRadarrInstances(ctx, client, instances)
		data, err := json.MarshalIndent(enriched, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// SonarrServicesResourceHandler returns a resource handler that lists all configured
// Sonarr instances enriched with per-instance quality profiles and root folders.
func SonarrServicesResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		instances, _, err := client.ServiceAPI.ServiceSonarrGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		enriched := enrichSonarrInstances(ctx, client, instances)
		data, err := json.MarshalIndent(enriched, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// SystemAboutResourceHandler returns a resource handler for Seerr system information.
func SystemAboutResourceHandler() server.ResourceHandlerFunc {
	return func(ctx context.Context, req mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		client := newAPIClientWithKey(apiKeyFromContext(ctx))
		res, _, err := client.SettingsAPI.SettingsAboutGet(ctx).Execute()
		if err != nil {
			return nil, err
		}
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return nil, err
		}
		return textResource(req.Params.URI, data), nil
	}
}

// enrichRadarrInstances fetches per-instance details for each Radarr entry concurrently
// and returns the enriched list. Instances without an ID fall back to the basic settings.
func enrichRadarrInstances(ctx context.Context, client *api.APIClient, instances []api.RadarrSettings) []interface{} {
	type detailResult struct {
		idx    int
		detail interface{}
	}
	ch := make(chan detailResult, len(instances))
	var wg sync.WaitGroup
	for i, inst := range instances {
		if !inst.HasId() {
			continue
		}
		wg.Add(1)
		go func(i int, id float32) {
			defer wg.Done()
			detail, _, err := client.ServiceAPI.ServiceRadarrRadarrIdGet(ctx, id).Execute()
			if err == nil {
				ch <- detailResult{idx: i, detail: detail}
			}
		}(i, inst.GetId())
	}
	wg.Wait()
	close(ch)

	detailsMap := make(map[int]interface{})
	for r := range ch {
		detailsMap[r.idx] = r.detail
	}

	result := make([]interface{}, len(instances))
	for i, inst := range instances {
		if d, ok := detailsMap[i]; ok {
			result[i] = d
		} else {
			result[i] = inst
		}
	}
	return result
}

// enrichSonarrInstances fetches per-instance details for each Sonarr entry concurrently
// and returns the enriched list. Instances without an ID fall back to the basic settings.
func enrichSonarrInstances(ctx context.Context, client *api.APIClient, instances []api.SonarrSettings) []interface{} {
	type detailResult struct {
		idx    int
		detail interface{}
	}
	ch := make(chan detailResult, len(instances))
	var wg sync.WaitGroup
	for i, inst := range instances {
		if !inst.HasId() {
			continue
		}
		wg.Add(1)
		go func(i int, id float32) {
			defer wg.Done()
			detail, _, err := client.ServiceAPI.ServiceSonarrSonarrIdGet(ctx, id).Execute()
			if err == nil {
				ch <- detailResult{idx: i, detail: detail}
			}
		}(i, inst.GetId())
	}
	wg.Wait()
	close(ch)

	detailsMap := make(map[int]interface{})
	for r := range ch {
		detailsMap[r.idx] = r.detail
	}

	result := make([]interface{}, len(instances))
	for i, inst := range instances {
		if d, ok := detailsMap[i]; ok {
			result[i] = d
		} else {
			result[i] = inst
		}
	}
	return result
}
