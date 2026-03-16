package tests

import (
	"net/http"
	"testing"
	"time"

	"seerr-cli/cmd/apiutil"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeServerURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "bare host",
			input: "https://host",
			want:  "https://host/api/v1",
		},
		{
			name:  "trailing slash",
			input: "https://host/",
			want:  "https://host/api/v1",
		},
		{
			name:  "multiple trailing slashes",
			input: "https://host///",
			want:  "https://host/api/v1",
		},
		{
			name:  "already has api/v1",
			input: "https://host/api/v1",
			want:  "https://host/api/v1",
		},
		{
			name:  "api/v1 with trailing slash",
			input: "https://host/api/v1/",
			want:  "https://host/api/v1",
		},
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "only slashes",
			input: "///",
			want:  "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, apiutil.NormalizeServerURL(tc.input))
		})
	}
}

func TestDefaultHTTPClientHasTimeout(t *testing.T) {
	// Verify the default HTTP client carries a 30 s timeout so requests cannot hang indefinitely.
	client := apiutil.NewAPIClientWithKeyAndTransport("", nil)
	cfg := client.GetConfig()
	require.NotNil(t, cfg.HTTPClient)
	assert.Equal(t, 30*time.Second, cfg.HTTPClient.Timeout)
}

func TestCustomTransportAlsoGetsTimeout(t *testing.T) {
	// Verify a custom transport still gets wrapped in a client with the default timeout.
	transport := &http.Transport{}
	client := apiutil.NewAPIClientWithKeyAndTransport("", transport)
	cfg := client.GetConfig()
	require.NotNil(t, cfg.HTTPClient)
	assert.Equal(t, 30*time.Second, cfg.HTTPClient.Timeout)
	assert.Equal(t, transport, cfg.HTTPClient.Transport)
}
