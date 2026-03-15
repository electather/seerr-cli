package tests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestCollectionCommands(t *testing.T) {
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
			name:           "get collection",
			args:           []string{"collection", "get", "537982"},
			mockResponse:   `{"id": 537982, "name": "Dune Collection"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"name": "Dune Collection"`,
			expectedPath:   "/api/v1/collection/537982",
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

			apiutil.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEERR_SERVER", server.URL)
			defer func() {
				apiutil.OverrideServerURL = ""
				os.Unsetenv("SEERR_SERVER")
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
