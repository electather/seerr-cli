# Fix #93 & #94: WatchlistPost and RequestPost fail with "unknown field id"

## Context

Issues #93 and #94 report that the `watchlist_add` and `request_create` MCP tools fail with:
```
WatchlistPost failed: <response body> (HTTP json: unknown field "id")
RequestPost failed: <response body> (HTTP json: unknown field "id")
```

**Root cause**: The `User.UnmarshalJSON()` in `pkg/api/model_user.go` (line 1073–1084) uses `json.NewDecoder(...).DisallowUnknownFields()`. This strict decoder is called whenever a `User` struct is deserialized — including nested user objects (`requestedBy`, `modifiedBy`) inside Watchlist and MediaRequest responses. When the actual Overseerr API returns a user with a field not in the generated `User` struct (e.g., `id` inside a nested `settings` sub-object, or any other undocumented field), the strict decoder rejects it.

**Established pattern**: `seerrclient/search.go` already solves the same class of problem for search/discover by using `RawGet()` to bypass the generated client's strict JSON decoding entirely. This fix applies the same approach to the two broken POST endpoints.

---

## Implementation Plan

### 1. Add `RawPost` to `cmd/apiutil/client.go`

Mirror `RawGet` but for POST. Accepts `ctx`, `client *api.APIClient`, `path string`, and `body []byte` (JSON). Sets `Content-Type: application/json`. Returns raw response `([]byte, error)`. Returns an error for HTTP 4xx/5xx responses.

### 2. Add `RawPost`/`RawPostCtx` to `internal/seerrclient/client.go`

Following the `RawGet`/`RawGetCtx` pattern already in the file:
```go
func (c *Client) RawPost(path string, body []byte) ([]byte, error)
func (c *Client) RawPostCtx(ctx context.Context, path string, body []byte) ([]byte, error)
```
Both delegate to `apiutil.RawPost`.

### 3. Add `internal/seerrclient/watchlist.go` (new file)

```go
// WatchlistPostCtx adds a media item to the watchlist, bypassing the generated
// client's strict JSON decoding which rejects undocumented fields in nested User objects.
func (c *Client) WatchlistPostCtx(ctx context.Context, tmdbId int, title, mediaType string) ([]byte, error)
```
Constructs the request body as `{"tmdbId": <n>, "title": "<s>", "mediaType": "<s>"}`, calls `c.RawPostCtx(ctx, "/watchlist", body)`.

### 4. Add `internal/seerrclient/request.go` (new file)

```go
// RequestPostCtx creates a new media request, bypassing the generated
// client's strict JSON decoding which rejects undocumented fields in nested User objects.
func (c *Client) RequestPostCtx(ctx context.Context, mediaType string, mediaId int, is4k bool) ([]byte, error)
```
Constructs the request body as `{"mediaType": "<s>", "mediaId": <n>}` (with optional `"is4k": true`), calls `c.RawPostCtx(ctx, "/request", body)`.

### 5. Update `cmd/mcp/tools_watchlist.go`

In `WatchlistAddHandler`:
- Replace `newAPIClientWithKey` + `client.WatchlistAPI.WatchlistPost(...).Execute()` with `seerrclient.NewWithKey(apiKeyFromContext(callCtx)).WatchlistPostCtx(callCtx, tmdbId, title, mediaType)`
- The result is `[]byte` — return directly via `mcp.NewToolResultText(string(b))`

`WatchlistRemoveHandler` is unaffected (DELETE, no body, 204 response — no decode issue).

### 6. Update `cmd/mcp/tools_request.go`

In `RequestCreateHandler`:
- Replace `newAPIClientWithKey` + `client.RequestAPI.RequestPost(...).Execute()` with `seerrclient.NewWithKey(apiKeyFromContext(callCtx)).RequestPostCtx(callCtx, mediaType, mediaId, is4k)`
- Return raw bytes directly.

### 7. Add `tests/mcp_watchlist_test.go` (new file)

Test `WatchlistAddHandler` with a mock response that includes:
- Extra nested fields in `requestedBy` (simulating the actual Overseerr response that caused the bug)
- The `id` field at the top level

Verify the handler succeeds and the result contains the expected data.

### 8. Add `tests/mcp_request_test.go` (new file)

Test `RequestCreateHandler` with a mock response that includes:
- Nested `requestedBy` and `modifiedBy` with extra fields
- `id`, `createdAt`, `updatedAt` at top level

Verify the handler succeeds and the result contains the expected data.

---

## Critical Files

| File | Role |
|------|------|
| `cmd/apiutil/client.go` | Add `RawPost` function |
| `internal/seerrclient/client.go` | Add `RawPost`/`RawPostCtx` methods |
| `internal/seerrclient/watchlist.go` | New: `WatchlistPostCtx` |
| `internal/seerrclient/request.go` | New: `RequestPostCtx` |
| `cmd/mcp/tools_watchlist.go` | Update `WatchlistAddHandler` |
| `cmd/mcp/tools_request.go` | Update `RequestCreateHandler` |
| `tests/mcp_watchlist_test.go` | New: MCP watchlist add test |
| `tests/mcp_request_test.go` | New: MCP request create test |

## Reuse

- `apiutil.RawGet` (`cmd/apiutil/client.go:43`) — template for `RawPost`
- `seerrclient.RawGet/RawGetCtx` (`internal/seerrclient/client.go:63-70`) — template for `RawPost/RawPostCtx`
- `seerrclient.SearchMultiCtx` (`internal/seerrclient/search.go:25`) — template for new wrapper methods
- `newMCPTestServer`/`callTool`/`resultText` (`tests/mcp_serve_test.go:19-44`) — test helpers to reuse
- `apiToolError` (`cmd/mcp/mcp.go:55`) — reuse for error handling in updated handlers

---

## Verification

```bash
# Run all tests
go test -v ./tests/ -run TestMCPWatchlistAdd
go test -v ./tests/ -run TestMCPRequestCreate
go test -v ./...

# Build to confirm no compile errors
go build
```

The new tests specifically use a mock response with extra/undocumented fields in nested user objects to confirm the bug is fixed. The tests should fail before the fix and pass after.
