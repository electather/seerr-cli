package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/tv"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestTvCommands(t *testing.T) {
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
			name:           "get tv",
			args:           []string{"tv", "get", "1396"},
			mockResponse:   `{"id": 1396, "name": "Breaking Bad"}`,
			expectedOutput: `"name": "Breaking Bad"`,
			expectedPath:   "/api/v1/tv/1396",
		},
		{
			name:           "season",
			args:           []string{"tv", "season", "1396", "1"},
			mockResponse:   `{"id": 1, "seasonNumber": 1, "name": "Season 1"}`,
			expectedOutput: `"seasonNumber": 1`,
			expectedPath:   "/api/v1/tv/1396/season/1",
		},
		{
			name:           "recommendations",
			args:           []string{"tv", "recommendations", "1396"},
			mockResponse:   `{"results": [{"id": 1, "name": "Better Call Saul", "mediaType": "tv"}]}`,
			expectedOutput: `"name": "Better Call Saul"`,
			expectedPath:   "/api/v1/tv/1396/recommendations",
		},
		{
			name:           "similar",
			args:           []string{"tv", "similar", "1396"},
			mockResponse:   `{"results": [{"id": 2, "name": "Ozark", "mediaType": "tv"}]}`,
			expectedOutput: `"name": "Ozark"`,
			expectedPath:   "/api/v1/tv/1396/similar",
		},
		{
			name:           "ratings",
			args:           []string{"tv", "ratings", "1396"},
			mockResponse:   `{"title": "Breaking Bad", "criticsScore": 96}`,
			expectedOutput: `"criticsScore": 96`,
			expectedPath:   "/api/v1/tv/1396/ratings",
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

			tv.OverrideServerURL = server.URL + "/api/v1"

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
