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
	"seer-cli/cmd/media"
)

func TestMediaCommands(t *testing.T) {
	viper.Set("server", "http://localhost:8080")
	viper.Set("api_key", "test-api-key")

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
			name:           "list media",
			args:           []string{"media", "list"},
			mockResponse:   `{"pageInfo": {"pages": 1, "pageSize": 20, "results": 2, "page": 1}, "results": [{"id": 1, "mediaType": "movie"}]}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"mediaType": "movie"`,
			expectedPath:   "/api/v1/media",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "list media with filter",
			args:           []string{"media", "list", "--filter", "available"},
			mockResponse:   `{"pageInfo": {"pages": 1, "pageSize": 20, "results": 1, "page": 1}, "results": [{"id": 2, "mediaType": "tv"}]}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"mediaType": "tv"`,
			expectedPath:   "/api/v1/media",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "delete media",
			args:           []string{"media", "delete", "42"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/media/42",
			expectedMethod: http.MethodDelete,
		},
		{
			name:           "delete media file",
			args:           []string{"media", "delete-file", "42"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/media/42/file",
			expectedMethod: http.MethodDelete,
		},
		{
			name:           "update media status",
			args:           []string{"media", "status", "42", "available"},
			mockResponse:   `{"id": 42, "mediaType": "movie", "status": 5}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"mediaType": "movie"`,
			expectedPath:   "/api/v1/media/42/available",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "get watch data",
			args:           []string{"media", "watch-data", "42"},
			mockResponse:   `{"data": {"playCount7Days": 3, "playCount30Days": 10, "playCount": 25, "users": []}}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"playCount7Days": 3`,
			expectedPath:   "/api/v1/media/42/watch_data",
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

			media.OverrideServerURL = server.URL + "/api/v1"

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
