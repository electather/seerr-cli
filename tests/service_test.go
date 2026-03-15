package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/service"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestServiceCommands(t *testing.T) {
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
			name:           "radarr list",
			args:           []string{"service", "radarr-list"},
			mockResponse:   `[{"id": 1, "name": "Radarr", "hostname": "radarr.local", "port": 7878, "apiKey": "key", "useSsl": false, "activeProfileId": 1, "activeProfileName": "HD", "activeDirectory": "/movies", "is4k": false, "minimumAvailability": "released", "isDefault": true}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"name": "Radarr"`,
			expectedPath:   "/api/v1/service/radarr",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "radarr get",
			args:           []string{"service", "radarr-get", "1"},
			mockResponse:   `{"server": {"id": 1, "name": "Radarr", "hostname": "radarr.local", "port": 7878, "apiKey": "key", "useSsl": false, "activeProfileId": 1, "activeProfileName": "HD", "activeDirectory": "/movies", "is4k": false, "minimumAvailability": "released", "isDefault": true}, "profiles": {"id": 1, "name": "HD-1080p"}}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"name": "Radarr"`,
			expectedPath:   "/api/v1/service/radarr/1",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "sonarr list",
			args:           []string{"service", "sonarr-list"},
			mockResponse:   `[{"id": 1, "name": "Sonarr", "hostname": "sonarr.local", "port": 8989, "apiKey": "key", "useSsl": false, "activeProfileId": 1, "activeProfileName": "HD", "activeDirectory": "/tv", "is4k": false, "enableSeasonFolders": true, "isDefault": true}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"name": "Sonarr"`,
			expectedPath:   "/api/v1/service/sonarr",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "sonarr get",
			args:           []string{"service", "sonarr-get", "1"},
			mockResponse:   `{"server": {"id": 1, "name": "Sonarr", "hostname": "sonarr.local", "port": 8989, "apiKey": "key", "useSsl": false, "activeProfileId": 1, "activeProfileName": "HD", "activeDirectory": "/tv", "is4k": false, "enableSeasonFolders": true, "isDefault": true}, "profiles": {"id": 1, "name": "HD-1080p"}}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"name": "Sonarr"`,
			expectedPath:   "/api/v1/service/sonarr/1",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "sonarr lookup",
			args:           []string{"service", "sonarr-lookup", "12345"},
			mockResponse:   `[{"title": "Breaking Bad", "tvdbId": 81189}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"title": "Breaking Bad"`,
			expectedPath:   "/api/v1/service/sonarr/lookup/12345",
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

			service.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				service.OverrideServerURL = ""
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
