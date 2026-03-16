package apiutil

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/viper"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

// NewAPIClient builds a standard API client using Viper config.
func NewAPIClient() (*api.APIClient, context.Context, bool) {
	return NewAPIClientWithTransport(nil), context.Background(), viper.GetBool("verbose")
}

// NewAPIClientWithTransport builds an API client with an optional custom HTTP transport.
// Pass nil to use the default transport.
func NewAPIClientWithTransport(transport http.RoundTripper) *api.APIClient {
	return NewAPIClientWithKeyAndTransport("", transport)
}

// NewAPIClientWithKey builds a client using apiKey, falling back to Viper when empty.
func NewAPIClientWithKey(apiKey string) *api.APIClient {
	return NewAPIClientWithKeyAndTransport(apiKey, nil)
}

// RawGet makes an authenticated GET request using the API client's configuration
// and returns the raw response body bytes. params may be nil for requests with
// no query parameters. Spaces in parameter values are encoded as %20, not +,
// to satisfy the Seerr API's strict URL-encoding requirement.
func RawGet(ctx context.Context, client *api.APIClient, path string, params url.Values) ([]byte, error) {
	cfg := client.GetConfig()
	serverURL, err := cfg.ServerURLWithContext(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("resolve server URL: %w", err)
	}
	u, err := url.Parse(serverURL + path)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}
	if len(params) > 0 {
		// url.Values.Encode() uses + for spaces; replace with %20 for Seerr compatibility.
		u.RawQuery = strings.ReplaceAll(params.Encode(), "+", "%20")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range cfg.DefaultHeader {
		req.Header.Set(k, v)
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}
	return body, nil
}

// NormalizeServerURL trims trailing slashes from raw and appends /api/v1 exactly
// once. Returns an empty string when raw is blank or all slashes.
func NormalizeServerURL(raw string) string {
	s := strings.TrimRight(raw, "/")
	if s == "" {
		return ""
	}
	if !strings.HasSuffix(s, "/api/v1") {
		s += "/api/v1"
	}
	return s
}

// NewAPIClientWithKeyAndTransport is the base constructor used by all other helpers.
func NewAPIClientWithKeyAndTransport(apiKey string, transport http.RoundTripper) *api.APIClient {
	configuration := api.NewConfiguration()
	serverURL := NormalizeServerURL(viper.GetString("seerr.server"))
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	key := apiKey
	if key == "" {
		key = viper.GetString("seerr.api_key")
	}
	if key != "" {
		configuration.AddDefaultHeader("X-Api-Key", key)
	}
	if transport != nil {
		configuration.HTTPClient = &http.Client{Transport: transport}
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}
	return api.NewAPIClient(configuration)
}
