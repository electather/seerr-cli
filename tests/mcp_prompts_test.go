package tests

import (
	"context"
	"testing"

	cmdmcp "seerr-cli/cmd/mcp"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// callPrompt invokes a prompt handler and returns the GetPromptResult.
func callPrompt(t *testing.T, handler func(context.Context, mcp.GetPromptRequest) (*mcp.GetPromptResult, error), args map[string]string) *mcp.GetPromptResult {
	t.Helper()
	req := mcp.GetPromptRequest{}
	req.Params.Arguments = args
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	return result
}

// promptText returns the text content of the first prompt message.
func promptText(t *testing.T, result *mcp.GetPromptResult) string {
	t.Helper()
	require.NotEmpty(t, result.Messages, "expected at least one message")
	content, ok := result.Messages[0].Content.(mcp.TextContent)
	require.True(t, ok, "expected TextContent in first message")
	return content.Text
}

func TestMCPRequestMediaPrompt(t *testing.T) {
	result := callPrompt(t, cmdmcp.RequestMediaPromptHandler(), map[string]string{
		"title": "Inception",
	})

	assert.NotEmpty(t, result.Description)
	text := promptText(t, result)
	assert.Contains(t, text, "Inception")
	assert.Contains(t, text, "search_multi")
	assert.Contains(t, text, "request_create")
	assert.Equal(t, mcp.RoleUser, result.Messages[0].Role)
}

func TestMCPRequestMediaPromptWith4K(t *testing.T) {
	result := callPrompt(t, cmdmcp.RequestMediaPromptHandler(), map[string]string{
		"title":      "Dune",
		"media_type": "movie",
		"quality":    "4k",
	})

	text := promptText(t, result)
	assert.Contains(t, text, "Dune")
	assert.Contains(t, text, "true") // is4k should be true
}

func TestMCPDiscoverContentPromptNoPreferences(t *testing.T) {
	result := callPrompt(t, cmdmcp.DiscoverContentPromptHandler(), nil)

	assert.NotEmpty(t, result.Description)
	text := promptText(t, result)
	assert.Contains(t, text, "search_trending")
}

func TestMCPDiscoverContentPromptWithPreferences(t *testing.T) {
	result := callPrompt(t, cmdmcp.DiscoverContentPromptHandler(), map[string]string{
		"preferences": "action thriller",
		"media_type":  "movie",
	})

	text := promptText(t, result)
	assert.Contains(t, text, "action thriller")
	assert.Contains(t, text, "seerr://genres")
}

func TestMCPManageRequestsPromptDefaultStatus(t *testing.T) {
	result := callPrompt(t, cmdmcp.ManageRequestsPromptHandler(), nil)

	text := promptText(t, result)
	assert.Contains(t, text, "request_count")
	assert.Contains(t, text, "pending")
	assert.Contains(t, text, "request_approve")
	assert.Contains(t, text, "request_decline")
}

func TestMCPManageRequestsPromptCustomStatus(t *testing.T) {
	result := callPrompt(t, cmdmcp.ManageRequestsPromptHandler(), map[string]string{
		"status": "failed",
	})

	text := promptText(t, result)
	assert.Contains(t, text, "failed")
}

func TestMCPReportIssuePrompt(t *testing.T) {
	result := callPrompt(t, cmdmcp.ReportIssuePromptHandler(), map[string]string{
		"media_title": "Breaking Bad",
	})

	assert.NotEmpty(t, result.Description)
	text := promptText(t, result)
	assert.Contains(t, text, "Breaking Bad")
	assert.Contains(t, text, "search_multi")
	assert.Contains(t, text, "issue_create")
}

func TestMCPReportIssuePromptWithDescription(t *testing.T) {
	result := callPrompt(t, cmdmcp.ReportIssuePromptHandler(), map[string]string{
		"media_title": "The Wire",
		"description": "Audio is out of sync",
	})

	text := promptText(t, result)
	assert.Contains(t, text, "Audio is out of sync")
}

func TestMCPMyDashboardPrompt(t *testing.T) {
	result := callPrompt(t, cmdmcp.MyDashboardPromptHandler(), nil)

	assert.NotEmpty(t, result.Description)
	text := promptText(t, result)
	assert.Contains(t, text, "auth_me")
	assert.Contains(t, text, "users_quota")
}

func TestMCPAdminOverviewPrompt(t *testing.T) {
	result := callPrompt(t, cmdmcp.AdminOverviewPromptHandler(), nil)

	assert.NotEmpty(t, result.Description)
	text := promptText(t, result)
	assert.Contains(t, text, "status_system")
	assert.Contains(t, text, "request_count")
	assert.Contains(t, text, "issue_count")
	assert.Contains(t, text, "settings_jobs_list")
}
