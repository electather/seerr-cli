package tests

import (
	"net/http"
	"testing"
	"time"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestHTTPServerHasNonZeroTimeouts(t *testing.T) {
	// Verify the HTTP server is created with non-zero timeout values to prevent
	// resource exhaustion from slow or stuck connections.
	srv := cmdmcp.NewHTTPServer(":8811", http.NewServeMux())
	assert.NotZero(t, srv.ReadHeaderTimeout)
	assert.NotZero(t, srv.ReadTimeout)
	assert.NotZero(t, srv.WriteTimeout)
	assert.NotZero(t, srv.IdleTimeout)
}

func TestHTTPServerTimeoutValues(t *testing.T) {
	// Verify the HTTP server timeout values match the documented safe defaults.
	srv := cmdmcp.NewHTTPServer(":8811", http.NewServeMux())
	assert.Equal(t, 5*time.Second, srv.ReadHeaderTimeout)
	assert.Equal(t, 15*time.Second, srv.ReadTimeout)
	assert.Equal(t, 30*time.Second, srv.WriteTimeout)
	assert.Equal(t, 60*time.Second, srv.IdleTimeout)
}

func TestHTTPServerAddrAndHandler(t *testing.T) {
	// Verify the Addr and Handler are set correctly.
	mux := http.NewServeMux()
	srv := cmdmcp.NewHTTPServer(":9999", mux)
	assert.Equal(t, ":9999", srv.Addr)
	assert.NotNil(t, srv.Handler)
}
