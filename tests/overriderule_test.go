package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/overriderule"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestOverrideRuleCommands(t *testing.T) {
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
			name:           "list override rules",
			args:           []string{"overriderule", "list"},
			mockResponse:   `[{"id": 1, "users": null, "genre": null, "language": null, "keywords": null, "profileId": null, "rootFolder": null, "tags": null, "radarrServiceId": null, "sonarrServiceId": null}]`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"id"`,
			expectedPath:   "/api/v1/overrideRule",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "create override rule",
			args:           []string{"overriderule", "create"},
			mockResponse:   `{"id": 2, "users": null, "genre": null, "language": null, "keywords": null, "profileId": null, "rootFolder": null, "tags": null, "radarrServiceId": null, "sonarrServiceId": null}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"id"`,
			expectedPath:   "/api/v1/overrideRule",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "update override rule",
			args:           []string{"overriderule", "update", "1"},
			mockResponse:   `{"id": 1, "users": null, "genre": null, "language": null, "keywords": null, "profileId": null, "rootFolder": null, "tags": null, "radarrServiceId": null, "sonarrServiceId": null}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"id"`,
			expectedPath:   "/api/v1/overrideRule/1",
			expectedMethod: http.MethodPut,
		},
		{
			name:           "delete override rule",
			args:           []string{"overriderule", "delete", "1"},
			mockResponse:   `{"id": 1, "users": null, "genre": null, "language": null, "keywords": null, "profileId": null, "rootFolder": null, "tags": null, "radarrServiceId": null, "sonarrServiceId": null}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"id"`,
			expectedPath:   "/api/v1/overrideRule/1",
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

			overriderule.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				overriderule.OverrideServerURL = ""
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
