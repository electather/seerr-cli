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
	"seer-cli/cmd/person"
)

func TestPersonCommands(t *testing.T) {
	viper.Set("server", "http://localhost:8080")
	viper.Set("api_key", "test-api-key")

	tests := []struct {
		name           string
		args           []string
		mockResponse   string
		expectedOutput string
		expectedPath   string
	}{
		{
			name:           "get",
			args:           []string{"person", "get", "287"},
			mockResponse:   `{"id": 287, "name": "Brad Pitt"}`,
			expectedOutput: `"name": "Brad Pitt"`,
			expectedPath:   "/api/v1/person/287",
		},
		{
			name:           "combined-credits",
			args:           []string{"person", "combined-credits", "287"},
			mockResponse:   `{"cast": [{"id": 550, "title": "Fight Club", "mediaType": "movie"}], "crew": []}`,
			expectedOutput: `"title": "Fight Club"`,
			expectedPath:   "/api/v1/person/287/combined_credits",
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

			person.OverrideServerURL = server.URL + "/api/v1"

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
