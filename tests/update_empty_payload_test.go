package tests

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"seerr-cli/cmd"
	"seerr-cli/cmd/apiutil"
	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// resetCmdFlags resets the Changed state of all flags on the named subcommand.
// Cobra does not reset Changed between Execute() calls, so this helper prevents
// flag state from leaking between tests that share the global command tree.
func resetCmdFlags(path ...string) {
	if c, _, _ := cmd.RootCmd.Find(path); c != nil {
		c.Flags().VisitAll(func(f *pflag.Flag) { f.Changed = false })
	}
}

// TestUsersUpdateEmptyPayload verifies that users update returns an error when
// no optional fields are provided so that the API is never called with an empty body.
func TestUsersUpdateEmptyPayload(t *testing.T) {
	resetCmdFlags("users", "update")

	// Point at a dummy server; it must not receive any requests.
	apiutil.OverrideServerURL = "http://127.0.0.1:0/api/v1"
	viper.Set("seerr.server", "http://127.0.0.1:0")
	defer func() { apiutil.OverrideServerURL = "" }()

	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)
	cmd.RootCmd.SetArgs([]string{"users", "update", "1"})

	err := cmd.RootCmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one field must be provided")
}

// TestRequestUpdateEmptyPayload verifies that request update returns an error
// when only the required --media-type flag is provided and no optional fields
// are set, preventing a no-op API call.
func TestRequestUpdateEmptyPayload(t *testing.T) {
	resetCmdFlags("request", "update")

	apiutil.OverrideServerURL = "http://127.0.0.1:0/api/v1"
	viper.Set("seerr.server", "http://127.0.0.1:0")
	defer func() { apiutil.OverrideServerURL = "" }()

	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)
	cmd.RootCmd.SetArgs([]string{"request", "update", "42", "--media-type", "movie"})

	err := cmd.RootCmd.Execute()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one field must be provided")
}

// TestMCPUsersUpdateEmptyPayload verifies that the MCP users_update tool
// returns an error result when only userId is provided and no update fields
// are set.
func TestMCPUsersUpdateEmptyPayload(t *testing.T) {
	handler := cmdmcp.UsersUpdateHandler()
	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{
		"userId": float64(1),
	}

	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.True(t, result.IsError, "expected IsError=true for empty payload")
	require.NotEmpty(t, result.Content)
	textContent, ok := result.Content[0].(mcp.TextContent)
	require.True(t, ok)
	assert.True(t, strings.Contains(textContent.Text, "at least one field must be provided"))
}
