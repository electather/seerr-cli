package tests

import (
	"net/http"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestMCPAuthMeHandler(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/auth/me" {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"id":1,"email":"admin@example.com","username":"admin","permissions":2}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	result := callTool(t, cmdmcp.AuthMeHandler(), nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"id"`)
	assert.Contains(t, text, `"email"`)
	assert.Contains(t, text, `"admin@example.com"`)
}

func TestMCPRequestRetryHandler(t *testing.T) {
	_, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/request/42/retry" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"id":42,"status":1,"type":"movie"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()

	result := callTool(t, cmdmcp.RequestRetryHandler(), map[string]any{"requestId": "42"})
	text := resultText(t, result)

	assert.Contains(t, text, `"id"`)
}
