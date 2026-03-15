package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"seer-cli/cmd"
	"seer-cli/cmd/blocklist"
)

func TestBlocklistCommands(t *testing.T) {
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
			name:           "list blocklist",
			args:           []string{"blocklist", "list"},
			mockResponse:   `{"pageInfo": {"pages": 1, "pageSize": 20, "results": 1, "page": 1}, "results": [{"tmdbId": 123, "title": "Dune"}]}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"title": "Dune"`,
			expectedPath:   "/api/v1/blocklist",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "list blocklist with filter",
			args:           []string{"blocklist", "list", "--filter", "all"},
			mockResponse:   `{"pageInfo": {"pages": 1, "pageSize": 20, "results": 0, "page": 1}, "results": []}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"results"`,
			expectedPath:   "/api/v1/blocklist",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "get blocklist item",
			args:           []string{"blocklist", "get", "123"},
			mockResponse:   `{"tmdbId": 123, "title": "Dune"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"tmdbId": 123`,
			expectedPath:   "/api/v1/blocklist/123",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "delete blocklist item",
			args:           []string{"blocklist", "delete", "123"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/blocklist/123",
			expectedMethod: http.MethodDelete,
		},
		{
			name:           "add to blocklist",
			args:           []string{"blocklist", "add", "--tmdb-id", "456"},
			mockResponse:   "",
			mockStatus:     http.StatusCreated,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/blocklist",
			expectedMethod: http.MethodPost,
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

			blocklist.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				blocklist.OverrideServerURL = ""
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
