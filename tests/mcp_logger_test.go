package tests

import (
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSafeLogPath(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		routeToken  string
		multiTenant bool
		want        string
	}{
		{
			name: "plain mcp path unchanged",
			path: "/mcp",
			want: "/mcp",
		},
		{
			name: "plain mcp sse path unchanged",
			path: "/mcp/sse",
			want: "/mcp/sse",
		},
		{
			name:       "route token in prefix is redacted",
			path:       "/abc123/mcp",
			routeToken: "abc123",
			want:       "/{redacted}/mcp",
		},
		{
			name:       "route token sse path is redacted",
			path:       "/abc123/mcp/sse",
			routeToken: "abc123",
			want:       "/{redacted}/mcp/sse",
		},
		{
			name:       "route token exact match is redacted",
			path:       "/abc123",
			routeToken: "abc123",
			want:       "/{redacted}",
		},
		{
			name:       "unrelated path is unchanged in route token mode",
			path:       "/health",
			routeToken: "abc123",
			want:       "/health",
		},
		{
			name:        "multi-tenant api key in path is redacted",
			path:        "/user-api-key/mcp",
			multiTenant: true,
			want:        "/{tenant}/mcp",
		},
		{
			name:        "multi-tenant api key with sse suffix is redacted",
			path:        "/user-api-key/mcp/sse",
			multiTenant: true,
			want:        "/{tenant}/mcp/sse",
		},
		{
			name:        "root path unchanged in multi-tenant mode",
			path:        "/",
			multiTenant: true,
			want:        "/",
		},
		{
			name:        "single segment path unchanged in multi-tenant mode",
			path:        "/health",
			multiTenant: true,
			want:        "/health",
		},
		{
			name:        "no route token and no multi-tenant returns path unchanged",
			path:        "/unexpected/path",
			routeToken:  "",
			multiTenant: false,
			want:        "/unexpected/path",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cmdmcp.SafeLogPath(tc.path, tc.routeToken, tc.multiTenant)
			assert.Equal(t, tc.want, got)
			// Verify the raw token never appears in the output.
			if tc.routeToken != "" {
				assert.NotContains(t, got, tc.routeToken)
			}
		})
	}
}
