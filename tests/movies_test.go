package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"seer-cli/cmd"
	"seer-cli/cmd/movies"
)

func TestMovieCommands(t *testing.T) {
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
			name:           "get movie",
			args:           []string{"movies", "get", "603"},
			mockResponse:   `{"id": 603, "title": "The Matrix"}`,
			expectedOutput: `"title": "The Matrix"`,
			expectedPath:   "/api/v1/movie/603",
		},
		{
			name:           "recommendations",
			args:           []string{"movies", "recommendations", "603"},
			mockResponse:   `{"results": [{"id": 1, "title": "Recommended Movie", "mediaType": "movie"}]}`,
			expectedOutput: `"title": "Recommended Movie"`,
			expectedPath:   "/api/v1/movie/603/recommendations",
		},
		{
			name:           "similar",
			args:           []string{"movies", "similar", "603"},
			mockResponse:   `{"results": [{"id": 1, "title": "Similar Movie", "mediaType": "movie"}]}`,
			expectedOutput: `"title": "Similar Movie"`,
			expectedPath:   "/api/v1/movie/603/similar",
		},
		{
			name:           "ratings",
			args:           []string{"movies", "ratings", "603"},
			mockResponse:   `{"title": "The Matrix", "criticsScore": 88}`,
			expectedOutput: `"criticsScore": 88`,
			expectedPath:   "/api/v1/movie/603/ratings",
		},
		{
			name:           "ratings-combined",
			args:           []string{"movies", "ratings-combined", "603"},
			mockResponse:   `{"imdb": {"title": "The Matrix", "criticsScore": 8.7}}`,
			expectedOutput: `"criticsScore": 8.7`,
			expectedPath:   "/api/v1/movie/603/ratingscombined",
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

			movies.OverrideServerURL = server.URL + "/api/v1"

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
