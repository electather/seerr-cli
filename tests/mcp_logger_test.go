package tests

import (
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestSafeLogQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  string
	}{
		{
			name:  "empty query unchanged",
			query: "",
			want:  "",
		},
		{
			name:  "api_key value is redacted",
			query: "api_key=secret123",
			want:  "api_key={redacted}",
		},
		{
			name:  "api_key redacted while other params preserved",
			query: "api_key=secret&page=1",
			want:  "api_key={redacted}&page=1",
		},
		{
			name:  "query without api_key is unchanged",
			query: "page=1&limit=10",
			want:  "page=1&limit=10",
		},
		{
			name:  "api_key in middle of query is redacted",
			query: "page=1&api_key=mysecret&limit=10",
			want:  "page=1&api_key={redacted}&limit=10",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cmdmcp.SafeLogQuery(tc.query)
			assert.Equal(t, tc.want, got)
		})
	}
}
