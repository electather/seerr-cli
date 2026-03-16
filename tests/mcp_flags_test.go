package tests

import (
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/stretchr/testify/assert"
)

func TestRouteTokenFlagNotRegistered(t *testing.T) {
	// --route-token has been removed; ensure it is not present in ServeFlags.
	for _, f := range cmdmcp.ServeFlags {
		assert.NotEqual(t, "route-token", f.Name, "--route-token must not appear in ServeFlags")
	}
}

func TestMultiTenantFlagNotRegistered(t *testing.T) {
	// --multi-tenant has been removed; ensure it is not present in ServeFlags.
	for _, f := range cmdmcp.ServeFlags {
		assert.NotEqual(t, "multi-tenant", f.Name, "--multi-tenant must not appear in ServeFlags")
	}
}

func TestAllowAPIKeyQueryParamFlagRegistered(t *testing.T) {
	// --allow-api-key-query-param must be present in ServeFlags.
	var found bool
	for _, f := range cmdmcp.ServeFlags {
		if f.Name == "allow-api-key-query-param" {
			found = true
			break
		}
	}
	assert.True(t, found, "--allow-api-key-query-param must appear in ServeFlags")
}
