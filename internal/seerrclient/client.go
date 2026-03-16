// Package seerrclient provides a clean wrapper around the auto-generated
// pkg/api client, hiding generated-code quirks like float32 IDs, broken
// union-type deserialization, and Seerr-specific URL encoding requirements.
package seerrclient

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/viper"
)

// Client wraps the generated API client and exposes clean Go types to callers.
type Client struct {
	api *api.APIClient
	ctx context.Context
}

// encodingRoundTripper replaces '+' with '%20' in query parameters. This is
// required because Seerr rejects the standard application/x-www-form-urlencoded
// plus-sign encoding for spaces.
type encodingRoundTripper struct {
	proxied http.RoundTripper
}

func (ert *encodingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.RawQuery = strings.ReplaceAll(req.URL.RawQuery, "+", "%20")
	return ert.proxied.RoundTrip(req)
}

// New creates a Client using Viper configuration (server URL and API key).
func New() *Client {
	return NewWithKey("")
}

// NewWithKey creates a Client with an explicit API key, falling back to Viper
// when key is empty. Used by MCP handlers that receive per-request API keys.
func NewWithKey(apiKey string) *Client {
	transport := &encodingRoundTripper{proxied: http.DefaultTransport}
	client := apiutil.NewAPIClientWithKeyAndTransport(apiKey, transport)
	return &Client{api: client, ctx: context.Background()}
}

// Ctx returns the context attached to this client.
func (c *Client) Ctx() context.Context {
	return c.ctx
}

// Unwrap returns the underlying generated API client for endpoints that have
// no wrapper method yet. Callers are responsible for any float32 casts.
func (c *Client) Unwrap() *api.APIClient {
	return c.api
}

// RawGet performs an authenticated GET request and returns the raw response
// body. Spaces in parameter values are encoded as %20 (not +) to satisfy
// Seerr's strict URL-encoding requirement.
func (c *Client) RawGet(path string, params url.Values) ([]byte, error) {
	return apiutil.RawGet(c.ctx, c.api, path, params)
}

// RawGetCtx performs a context-aware authenticated GET request.
func (c *Client) RawGetCtx(ctx context.Context, path string, params url.Values) ([]byte, error) {
	return apiutil.RawGet(ctx, c.api, path, params)
}

// Verbose returns whether verbose mode is enabled via Viper configuration.
func Verbose() bool {
	return viper.GetBool("verbose")
}
