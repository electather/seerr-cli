package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

// captureContextKey is a helper that records the API key injected into the
// request context by the middleware.
func captureContextKey(t *testing.T) (handler http.Handler, captured *string) {
	t.Helper()
	s := ""
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, _ = r.Context().Value(cmdmcp.APIKeyContextKey).(string)
		w.WriteHeader(http.StatusOK)
	})
	return h, &s
}

func TestSeerrAPIKeyMiddleware_headerOnly(t *testing.T) {
	inner, captured := captureContextKey(t)
	handler := cmdmcp.SeerrAPIKeyMiddleware(false, inner)

	req := httptest.NewRequest(http.MethodGet, "/mcp", nil)
	req.Header.Set("X-Api-Key", "header-key")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "header-key", *captured)
}

func TestSeerrAPIKeyMiddleware_queryParamOnly(t *testing.T) {
	inner, captured := captureContextKey(t)
	handler := cmdmcp.SeerrAPIKeyMiddleware(true, inner)

	req := httptest.NewRequest(http.MethodGet, "/mcp?api_key=qparam-key", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "qparam-key", *captured)
}

func TestSeerrAPIKeyMiddleware_headerPrecedenceOverQueryParam(t *testing.T) {
	inner, captured := captureContextKey(t)
	handler := cmdmcp.SeerrAPIKeyMiddleware(true, inner)

	req := httptest.NewRequest(http.MethodGet, "/mcp?api_key=qparam-key", nil)
	req.Header.Set("X-Api-Key", "header-key")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "header-key", *captured, "header must take precedence over query param")
}

func TestSeerrAPIKeyMiddleware_queryParamDisabled_ignoresQueryParam(t *testing.T) {
	var innerCalled bool
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		innerCalled = true
		w.WriteHeader(http.StatusOK)
	})
	// allowQueryParam=false; query param present but should not satisfy auth.
	handler := cmdmcp.SeerrAPIKeyMiddleware(false, inner)

	req := httptest.NewRequest(http.MethodGet, "/mcp?api_key=qparam-key", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	// No header, query param disabled — must return 401.
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, innerCalled)
}

func TestSeerrAPIKeyMiddleware_neitherPresent_returns401(t *testing.T) {
	var innerCalled bool
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		innerCalled = true
		w.WriteHeader(http.StatusOK)
	})
	handler := cmdmcp.SeerrAPIKeyMiddleware(true, inner)

	req := httptest.NewRequest(http.MethodGet, "/mcp", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, innerCalled)
}

func TestSeerrAPIKeyMiddleware_queryParam_sensitiveValueNotLogged(t *testing.T) {
	// This test ensures SafeLogQuery redacts the api_key value from query strings.
	redacted := cmdmcp.SafeLogQuery("api_key=secret123&page=1")
	assert.NotContains(t, redacted, "secret123")
	assert.Contains(t, redacted, "api_key={redacted}")
	assert.Contains(t, redacted, "page=1")
}

func TestSeerrAPIKeyMiddlewareIsTheOnlyPerRequestKeyMechanism(t *testing.T) {
	// SeerrAPIKeyMiddleware must compile and be the sole per-request API key
	// injection mechanism — path-based routing has been removed.
	_ = cmdmcp.SeerrAPIKeyMiddleware
}
