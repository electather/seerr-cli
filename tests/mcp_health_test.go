package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMCPHealthEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	cmdmcp.HealthCheckHandler(w, req)

	resp := w.Result()
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	body := w.Body.String()
	assert.Contains(t, body, `"status":"ok"`)
	assert.Contains(t, body, `"version"`)
}
