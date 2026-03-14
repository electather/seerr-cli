package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cmdmcp "seer-cli/cmd/mcp"
)

func newMCPTestServer(handler http.HandlerFunc) (*httptest.Server, func()) {
	ts := httptest.NewServer(handler)
	cmdmcp.OverrideServerURL = ts.URL + "/api/v1"
	return ts, func() {
		ts.Close()
		cmdmcp.OverrideServerURL = ""
	}
}

func callTool(t *testing.T, handler func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error), args map[string]any) *mcp.CallToolResult {
	t.Helper()
	req := mcp.CallToolRequest{}
	req.Params.Arguments = args
	result, err := handler(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	return result
}

func resultText(t *testing.T, result *mcp.CallToolResult) string {
	t.Helper()
	require.NotEmpty(t, result.Content)
	textContent, ok := result.Content[0].(mcp.TextContent)
	require.True(t, ok, "expected TextContent")
	return textContent.Text
}

func TestMCPStatusSystemHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/status" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"version":"1.0.0","commitTag":"abc123"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.StatusSystemHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"version"`)
	assert.Contains(t, text, `"1.0.0"`)
}

func TestMCPSearchMultiHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/search" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"page":1,"totalPages":1,"totalResults":1,"results":[]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.SearchMultiHandler(client, ctx)

	result := callTool(t, handler, map[string]any{"query": "batman"})
	text := resultText(t, result)

	assert.Contains(t, text, `"results"`)
}

func TestMCPSearchMultiHandlerMissingQuery(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.SearchMultiHandler(client, ctx)

	req := mcp.CallToolRequest{}
	req.Params.Arguments = map[string]any{}
	_, err := handler(context.Background(), req)
	assert.Error(t, err)
}

func TestMCPRequestListHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/request" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"pageInfo":{"pages":1,"pageSize":10,"results":0,"page":1},"results":[]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.RequestListHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"results"`)
}

func TestMCPRequestCountHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/request/count" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"total":5,"movie":3,"tv":2,"pending":1,"approved":4}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.RequestCountHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"total"`)
}

func TestMCPMoviesGetHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/v1/movie/") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":550,"title":"Fight Club","mediaType":"movie"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.MoviesGetHandler(client, ctx)

	result := callTool(t, handler, map[string]any{"movieId": float64(550)})
	text := resultText(t, result)

	assert.Contains(t, text, `"title"`)
	assert.Contains(t, text, `"Fight Club"`)
}

func TestMCPTVGetHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/v1/tv/") && !strings.Contains(r.URL.Path, "season") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":1399,"name":"Game of Thrones","mediaType":"tv"}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.TVGetHandler(client, ctx)

	result := callTool(t, handler, map[string]any{"tvId": float64(1399)})
	text := resultText(t, result)

	assert.Contains(t, text, `"name"`)
	assert.Contains(t, text, `"Game of Thrones"`)
}

func TestMCPIssueCountHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/issue/count" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"total":3,"video":1,"audio":1,"subtitle":0,"other":1}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.IssueCountHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"total"`)
}

func TestMCPUsersListHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/user" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"pageInfo":{"pages":1,"pageSize":10,"results":1,"page":1},"results":[{"id":1,"displayName":"admin"}]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.UsersListHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"results"`)
}

func TestMCPServiceRadarrListHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/service/radarr" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"id":1,"name":"Radarr","hostname":"localhost","port":7878,"apiKey":"testkey","useSsl":false,"activeProfileId":1,"activeProfileName":"Any","activeDirectory":"/movies","is4k":false,"minimumAvailability":"announced","isDefault":true}]`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.ServiceRadarrListHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	// The result is an array of RadarrSettings
	var parsed []map[string]interface{}
	err := json.Unmarshal([]byte(text), &parsed)
	require.NoError(t, err)
	assert.Len(t, parsed, 1)
	assert.Equal(t, "Radarr", parsed[0]["name"])
}

func TestMCPSettingsAboutHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/settings/about" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"version":"1.0.0","totalRequests":100}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.SettingsAboutHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"version"`)
}

func TestMCPBlocklistListHandler(t *testing.T) {
	ts, cleanup := newMCPTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/blocklist" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"pageInfo":{"pages":1,"pageSize":10,"results":0,"page":1},"results":[]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	})
	defer cleanup()
	_ = ts

	client, ctx := cmdmcp.NewAPIClientForTest()
	handler := cmdmcp.BlocklistListHandler(client, ctx)

	result := callTool(t, handler, nil)
	text := resultText(t, result)

	assert.Contains(t, text, `"results"`)
}
