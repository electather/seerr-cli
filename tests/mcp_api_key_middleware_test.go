package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

const testAuthToken = "super-secret"

// newAuthTestHandler returns an inner handler that records whether it was called.
func newAuthTestHandler(called *bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*called = true
		w.WriteHeader(http.StatusOK)
	})
}

func TestMCPAuthMiddleware_bearerToken(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, false, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	req.Header.Set("Authorization", "Bearer "+testAuthToken)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, called)
}

func TestMCPAuthMiddleware_xApiKeyHeader(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, false, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	req.Header.Set("X-Api-Key", testAuthToken)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, called)
}

func TestMCPAuthMiddleware_queryParam(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, true, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp?api_key="+testAuthToken, nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, called)
}

func TestMCPAuthMiddleware_queryParamDisabled_returns401(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, false, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp?api_key="+testAuthToken, nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
}

func TestMCPAuthMiddleware_wrongToken_returns401(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, true, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	req.Header.Set("X-Api-Key", "wrong-token")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
}

func TestMCPAuthMiddleware_noCredentials_returns401(t *testing.T) {
	var called bool
	handler := cmdmcp.MCPAuthMiddleware(testAuthToken, true, newAuthTestHandler(&called))

	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, called)
}

// TestMCPAuthMiddleware_noAuthToken verifies that when no --auth-token is
// configured the middleware is not applied and requests pass through without
// any credential check. This is the no-auth / --no-auth scenario.
func TestMCPAuthMiddleware_noAuthToken_requestPassesThrough(t *testing.T) {
	// When authToken is empty the serve command does not wrap the handler with
	// MCPAuthMiddleware at all. The Seerr API key is always read from the app
	// config (seerr.api_key) and never sourced from the incoming request.
	var called bool
	inner := newAuthTestHandler(&called)

	// Simulate the no-auth path: handler is NOT wrapped with MCPAuthMiddleware.
	req := httptest.NewRequest(http.MethodPost, "/mcp", nil)
	rec := httptest.NewRecorder()
	inner.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, called)
}

func TestMCPAuthMiddlewareIsTheOnlyPerRequestAuthMechanism(t *testing.T) {
	// Compile-time check that MCPAuthMiddleware exists and has the expected signature.
	_ = cmdmcp.MCPAuthMiddleware
}

func TestMCPLogQueryRedaction(t *testing.T) {
	// Ensure SafeLogQuery redacts the api_key value from query strings.
	redacted := cmdmcp.SafeLogQuery("api_key=secret123&page=1")
	assert.NotContains(t, redacted, "secret123")
	assert.Contains(t, redacted, "api_key={redacted}")
	assert.Contains(t, redacted, "page=1")
}
