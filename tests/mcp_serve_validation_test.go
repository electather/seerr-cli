package tests

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cmdmcp "seerr-cli/cmd/mcp"
)

func TestMCPServeFailsFastWithoutSeerrServer(t *testing.T) {
	original := viper.GetString("seerr.server")
	t.Cleanup(func() { viper.Set("seerr.server", original) })

	tests := []struct {
		name        string
		seerrServer string
		wantErr     bool
		errContains string
	}{
		{
			name:        "missing server returns error",
			seerrServer: "",
			wantErr:     true,
			errContains: "seerr.server",
		},
		{
			name:        "only slashes returns error",
			seerrServer: "///",
			wantErr:     true,
			errContains: "seerr.server",
		},
		{
			name:        "valid server passes validation",
			seerrServer: "http://localhost:5055",
			wantErr:     false,
		},
		{
			name:        "valid server with trailing slash passes validation",
			seerrServer: "http://localhost:5055/",
			wantErr:     false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			viper.Set("seerr.server", tc.seerrServer)
			err := cmdmcp.ValidateServeConfig()
			if tc.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.errContains)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
