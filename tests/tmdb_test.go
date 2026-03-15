package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/tmdb"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestTmdbCommands(t *testing.T) {
	viper.Set("seer.server", "http://localhost:8080")
	viper.Set("seer.api_key", "test-api-key")

	tests := []struct {
		name           string
		args           []string
		mockResponse   string
		mockStatus     int
		expectedOutput string
		expectedPath   string
		expectedMethod string
	}{
		{
			name:           "regions",
			args:           []string{"tmdb", "regions"},
			mockResponse:   `[{"iso_3166_1": "US", "english_name": "United States of America"}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"iso_3166_1"`,
			expectedPath:   "/api/v1/regions",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "languages",
			args:           []string{"tmdb", "languages"},
			mockResponse:   `[{"iso_639_1": "en", "english_name": "English", "name": "English"}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"english_name"`,
			expectedPath:   "/api/v1/languages",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "backdrops",
			args:           []string{"tmdb", "backdrops"},
			mockResponse:   `["/path/to/backdrop.jpg"]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `backdrop`,
			expectedPath:   "/api/v1/backdrops",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "genres-movie",
			args:           []string{"tmdb", "genres-movie"},
			mockResponse:   `[{"id": 10751, "name": "Family"}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"Family"`,
			expectedPath:   "/api/v1/genres/movie",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "genres-tv",
			args:           []string{"tmdb", "genres-tv"},
			mockResponse:   `[{"id": 18, "name": "Drama"}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"Drama"`,
			expectedPath:   "/api/v1/genres/tv",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "studio",
			args:           []string{"tmdb", "studio", "2"},
			mockResponse:   `{"id": 2, "name": "Walt Disney Pictures"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"Walt Disney Pictures"`,
			expectedPath:   "/api/v1/studio/2",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "network",
			args:           []string{"tmdb", "network", "1"},
			mockResponse:   `{"id": 1, "name": "Fuji TV"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"Fuji TV"`,
			expectedPath:   "/api/v1/network/1",
			expectedMethod: http.MethodGet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedPath, r.URL.Path)
				assert.Equal(t, tt.expectedMethod, r.Method)
				assert.Equal(t, "test-api-key", r.Header.Get("X-Api-Key"))

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tt.mockStatus)
				if tt.mockResponse != "" {
					fmt.Fprintln(w, tt.mockResponse)
				}
			}))
			defer server.Close()

			tmdb.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				tmdb.OverrideServerURL = ""
				os.Unsetenv("SEER_SERVER")
			}()

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
