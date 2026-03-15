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
	"seer-cli/cmd/issue"
)

func TestIssueCommands(t *testing.T) {
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
			name:           "list issues",
			args:           []string{"issue", "list"},
			mockResponse:   `{"pageInfo": {"pages": 1, "pageSize": 20, "results": 1, "page": 1}, "results": [{"id": 1, "status": 1}]}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"results"`,
			expectedPath:   "/api/v1/issue",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "get issue count",
			args:           []string{"issue", "count"},
			mockResponse:   `{"total": 5, "open": 3, "closed": 2, "video": 1, "audio": 2, "subtitles": 1, "others": 1}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"total"`,
			expectedPath:   "/api/v1/issue/count",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "get issue",
			args:           []string{"issue", "get", "1"},
			mockResponse:   `{"id": 1, "status": 1, "issueType": 1}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"issueType"`,
			expectedPath:   "/api/v1/issue/1",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "create issue",
			args:           []string{"issue", "create", "1", "--message", "Audio is broken", "--media-id", "42"},
			mockResponse:   `{"id": 2, "status": 1, "issueType": 1}`,
			mockStatus:     http.StatusCreated,
			expectedOutput: `"id"`,
			expectedPath:   "/api/v1/issue",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "delete issue",
			args:           []string{"issue", "delete", "1"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/issue/1",
			expectedMethod: http.MethodDelete,
		},
		{
			name:           "update issue status",
			args:           []string{"issue", "update-status", "1", "resolved"},
			mockResponse:   `{"id": 1, "status": 2}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"status"`,
			expectedPath:   "/api/v1/issue/1/resolved",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "add comment",
			args:           []string{"issue", "comment", "1", "--message", "Looking into this"},
			mockResponse:   `{"id": 1, "comments": [{"id": 42, "message": "Looking into this"}]}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"comments"`,
			expectedPath:   "/api/v1/issue/1/comment",
			expectedMethod: http.MethodPost,
		},
		{
			name:           "get comment",
			args:           []string{"issue", "get-comment", "42"},
			mockResponse:   `{"id": 42, "message": "Looking into this"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"message"`,
			expectedPath:   "/api/v1/issueComment/42",
			expectedMethod: http.MethodGet,
		},
		{
			name:           "update comment",
			args:           []string{"issue", "update-comment", "42", "--message", "Updated message"},
			mockResponse:   `{"id": 42, "message": "Updated message"}`,
			mockStatus:     http.StatusOK,
			expectedOutput: `"Updated message"`,
			expectedPath:   "/api/v1/issueComment/42",
			expectedMethod: http.MethodPut,
		},
		{
			name:           "delete comment",
			args:           []string{"issue", "delete-comment", "42"},
			mockResponse:   "",
			mockStatus:     http.StatusNoContent,
			expectedOutput: `"status": "ok"`,
			expectedPath:   "/api/v1/issueComment/42",
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

			issue.OverrideServerURL = server.URL + "/api/v1"
			os.Setenv("SEER_SERVER", server.URL)
			defer func() {
				issue.OverrideServerURL = ""
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
