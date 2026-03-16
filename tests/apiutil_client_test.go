package tests

import (
	"testing"

	"seerr-cli/cmd/apiutil"

	"github.com/stretchr/testify/assert"
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
