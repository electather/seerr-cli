package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"seerr-cli/cmd"
	"seerr-cli/cmd/search"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSearchCommands(t *testing.T) {
	viper.Set("seer.server", "http://localhost:8080")
	viper.Set("seer.api_key", "test-api-key")

	tests := []struct {
		name           string
		args           []string
		mockResponse   string
		expectedOutput string
		expectedPath   string
	}{
		{
			name:           "multi search",
			args:           []string{"search", "multi", "-q", "Inception"},
			mockResponse:   `{"results": [{"id": 1, "title": "Inception", "mediaType": "movie"}]}`,
			expectedOutput: `"title": "Inception"`,
			expectedPath:   "/api/v1/search",
		},
		{
			name:           "keyword search",
			args:           []string{"search", "keyword", "-q", "sci-fi"},
			mockResponse:   `{"results": [{"id": 1, "name": "sci-fi"}]}`,
			expectedOutput: `"name": "sci-fi"`,
			expectedPath:   "/api/v1/search/keyword",
		},
		{
			name:           "company search",
			args:           []string{"search", "company", "-q", "Warner"},
			mockResponse:   `{"results": [{"id": 1, "name": "Warner Bros."}]}`,
			expectedOutput: `"name": "Warner Bros."`,
			expectedPath:   "/api/v1/search/company",
		},
		{
			name:           "trending search",
			args:           []string{"search", "trending"},
			mockResponse:   `{"results": [{"id": 1, "title": "Trending Movie", "mediaType": "movie"}]}`,
			expectedOutput: `"title": "Trending Movie"`,
			expectedPath:   "/api/v1/discover/trending",
		},
		{
			name:           "discover movies",
			args:           []string{"search", "movies", "--genre", "18"},
			mockResponse:   `{"results": [{"id": 1, "title": "Drama Movie", "mediaType": "movie"}]}`,
			expectedOutput: `"title": "Drama Movie"`,
			expectedPath:   "/api/v1/discover/movies",
		},
		{
			name:           "discover tv",
			args:           []string{"search", "tv", "--network", "1"},
			mockResponse:   `{"results": [{"id": 1, "name": "TV Show"}]}`,
			expectedOutput: `"name": "TV Show"`,
			expectedPath:   "/api/v1/discover/tv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedPath, r.URL.Path)
				assert.Equal(t, "test-api-key", r.Header.Get("X-Api-Key"))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, tt.mockResponse)
			}))
			defer server.Close()

			search.OverrideServerURL = server.URL + "/api/v1"

			b := bytes.NewBufferString("")
			command := cmd.RootCmd
			command.SetOut(b)
			command.SetArgs(tt.args)

			err := command.Execute()
			assert.NoError(t, err)
			assert.Contains(t, b.String(), tt.expectedOutput)
		})
	}
}
