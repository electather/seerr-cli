package tests

import (
	"context"
	"net/http"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// callResource invokes a resource handler and returns the text of the first result.
func callResource(t *testing.T, handler func(context.Context, mcp.ReadResourceRequest) ([]mcp.ResourceContents, error), uri string) string {
	t.Helper()
	req := mcp.ReadResourceRequest{}
	req.Params.URI = uri
	contents, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.NotEmpty(t, contents)
	text, ok := contents[0].(mcp.TextResourceContents)
	require.True(t, ok, "expected TextResourceContents")
	return text.Text
}

func TestMCPGenresMovieResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/genres/movie" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"id":28,"name":"Action"},{"id":12,"name":"Adventure"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.GenresMovieResourceHandler(), "seerr://genres/movies")
	assert.Contains(t, text, `"Action"`)
	assert.Contains(t, text, `"Adventure"`)
}

func TestMCPGenresTVResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/genres/tv" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"id":10765,"name":"Sci-Fi & Fantasy"},{"id":18,"name":"Drama"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.GenresTVResourceHandler(), "seerr://genres/tv")
	assert.Contains(t, text, `"Drama"`)
}

func TestMCPLanguagesResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/languages" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"iso_639_1":"en","english_name":"English"},{"iso_639_1":"fr","english_name":"French"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.LanguagesResourceHandler(), "seerr://languages")
	assert.Contains(t, text, `"English"`)
	assert.Contains(t, text, `"French"`)
}

func TestMCPRegionsResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/regions" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`[{"iso_3166_1":"US","english_name":"United States"},{"iso_3166_1":"GB","english_name":"United Kingdom"}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.RegionsResourceHandler(), "seerr://regions")
	assert.Contains(t, text, `"United States"`)
}

func TestMCPCertificationsMovieResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/certifications/movie" {
			w.Header().Set("Content-Type", "application/json")
			// The CertificationResponse struct wraps the map under a "certifications" key.
			w.Write([]byte(`{"certifications":{"US":[{"certification":"PG-13","meaning":"Parents strongly cautioned","order":3}]}}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.CertificationsMovieResourceHandler(), "seerr://certifications/movies")
	assert.Contains(t, text, `"PG-13"`)
}

func TestMCPCertificationsTVResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/certifications/tv" {
			w.Header().Set("Content-Type", "application/json")
			// The CertificationResponse struct wraps the map under a "certifications" key.
			w.Write([]byte(`{"certifications":{"US":[{"certification":"TV-MA","meaning":"Mature audiences","order":6}]}}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.CertificationsTVResourceHandler(), "seerr://certifications/tv")
	assert.Contains(t, text, `"TV-MA"`)
}

func TestMCPRadarrServicesResourceEnriched(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/service/radarr":
			w.Write([]byte(`[{"id":1,"name":"Radarr HD","hostname":"radarr","port":7878,"apiKey":"key","useSsl":false,"activeProfileId":1,"activeProfileName":"HD-1080p","activeDirectory":"/movies","is4k":false,"minimumAvailability":"released","isDefault":true}]`))
		case "/api/v1/service/radarr/1":
			w.Write([]byte(`{"server":{"id":1,"name":"Radarr HD","hostname":"radarr","port":7878,"apiKey":"key","useSsl":false,"activeProfileId":1,"activeProfileName":"HD-1080p","activeDirectory":"/movies","is4k":false,"minimumAvailability":"released","isDefault":true},"profiles":{"qualityProfiles":[{"id":1,"name":"HD-1080p"}],"rootFolders":[{"id":1,"path":"/movies"}]}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()

	text := callResource(t, cmdmcp.RadarrServicesResourceHandler(), "seerr://services/radarr")
	assert.Contains(t, text, `"profiles"`)
	assert.Contains(t, text, `"HD-1080p"`)
}

func TestMCPSonarrServicesResourceEnriched(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/v1/service/sonarr":
			w.Write([]byte(`[{"id":1,"name":"Sonarr","hostname":"sonarr","port":8989,"apiKey":"key","useSsl":false,"activeProfileId":1,"activeProfileName":"HD-1080p","activeDirectory":"/tv","is4k":false,"enableSeasonFolders":true,"isDefault":true}]`))
		case "/api/v1/service/sonarr/1":
			w.Write([]byte(`{"server":{"id":1,"name":"Sonarr","hostname":"sonarr","port":8989,"apiKey":"key","useSsl":false,"activeProfileId":1,"activeProfileName":"HD-1080p","activeDirectory":"/tv","is4k":false,"enableSeasonFolders":true,"isDefault":true},"profiles":{"qualityProfiles":[{"id":1,"name":"HD-1080p"}],"rootFolders":[{"id":1,"path":"/tv"}]}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	defer cleanup()

	text := callResource(t, cmdmcp.SonarrServicesResourceHandler(), "seerr://services/sonarr")
	assert.Contains(t, text, `"profiles"`)
}

func TestMCPSystemAboutResource(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/settings/about" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"version":"2.0.0","totalRequests":42}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	text := callResource(t, cmdmcp.SystemAboutResourceHandler(), "seerr://system/about")
	assert.Contains(t, text, `"2.0.0"`)
	assert.Contains(t, text, `"totalRequests"`)
}
