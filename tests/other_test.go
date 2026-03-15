package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/other"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestOtherCommands(t *testing.T) {
	viper.Set("seerr.server", "http://localhost:8080")
	viper.Set("seerr.api_key", "test-api-key")

	tests := []struct {
		name           string
		args           []string
		mockResponse   string
		expectedOutput string
		expectedPath   string
	}{
		{
			name:           "keyword",
			args:           []string{"other", "keyword", "1"},
			mockResponse:   `{"id": 1, "name": "superhero"}`,
			expectedOutput: `"name": "superhero"`,
			expectedPath:   "/api/v1/keyword/1",
		},
		{
			name:           "watchprovider-regions",
			args:           []string{"other", "watchprovider-regions"},
			mockResponse:   `[{"iso_3166_1": "US", "english_name": "United States"}]`,
			expectedOutput: `"english_name": "United States"`,
			expectedPath:   "/api/v1/watchproviders/regions",
		},
		{
			name:           "watchproviders-movies",
			args:           []string{"other", "watchproviders-movies", "--watch-region", "US"},
			mockResponse:   `[{"id": 8, "name": "Netflix"}]`,
			expectedOutput: `"name": "Netflix"`,
			expectedPath:   "/api/v1/watchproviders/movies",
		},
		{
			name:           "watchproviders-tv",
			args:           []string{"other", "watchproviders-tv", "--watch-region", "US"},
			mockResponse:   `[{"id": 8, "name": "Netflix"}]`,
			expectedOutput: `"name": "Netflix"`,
			expectedPath:   "/api/v1/watchproviders/tv",
		},
		{
			name:           "certifications-movie",
			args:           []string{"other", "certifications-movie"},
			mockResponse:   `{"certifications": {"US": [{"certification": "PG-13"}]}}`,
			expectedOutput: `"PG-13"`,
			expectedPath:   "/api/v1/certifications/movie",
		},
		{
			name:           "certifications-tv",
			args:           []string{"other", "certifications-tv"},
			mockResponse:   `{"certifications": {"US": [{"certification": "TV-MA"}]}}`,
			expectedOutput: `"TV-MA"`,
			expectedPath:   "/api/v1/certifications/tv",
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

			other.OverrideServerURL = server.URL + "/api/v1"

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
