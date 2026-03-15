package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/watchlist"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestWatchlistCommands(t *testing.T) {
	viper.Set("seerr.server", "http://localhost:8080")
	viper.Set("seerr.api_key", "test-api-key")

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
			name:           "add to watchlist",
			args:           []string{"watchlist", "add", "--tmdb-id", "123", "--media-type", "movie"},
			mockResponse:   `{"id": 1, "tmdbId": 123, "type": "movie", "title": "Dune"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"tmdbId": 123`,
			expectedPath:   "/api/v1/watchlist",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "delete from watchlist",
			args:           []string{"watchlist", "delete", "123"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/watchlist/123",
			expectedMethod: http.MethodDelete,
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

			watchlist.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				watchlist.OverrideServerURL = ""
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
