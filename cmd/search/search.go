package search

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	api "seer-cli/pkg/api"
)

// encodingRoundTripper is a custom RoundTripper that replaces '+' with '%20' in query parameters.
// It also removes problematic parameters from certain endpoints if needed.
type encodingRoundTripper struct {
	Proxied http.RoundTripper
}

func (ert *encodingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Special handling for /discover/trending: remove mediaType and timeWindow if they are defaults
	// because some server versions don't support them and fail with 400.
	if strings.Contains(req.URL.Path, "/discover/trending") {
		q := req.URL.Query()
		if q.Get("mediaType") == "all" {
			q.Del("mediaType")
		}
		if q.Get("timeWindow") == "day" {
			q.Del("timeWindow")
		}
		req.URL.RawQuery = q.Encode()
	}

	// Replace '+' with '%20' in the RawQuery
	req.URL.RawQuery = strings.ReplaceAll(req.URL.RawQuery, "+", "%20")

	return ert.Proxied.RoundTrip(req)
}

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

var Cmd = &cobra.Command{
	Use:   "search",
	Short: "Search for movies, TV shows, people, and more",
	Long:  `Search for various resources from the Seerr API including movies, TV shows, people, keywords, and companies.`,
}

func newAPIClient() (*api.APIClient, context.Context, bool) {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("seer.server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	if apiKey := viper.GetString("seer.api_key"); apiKey != "" {
		configuration.AddDefaultHeader("X-Api-Key", apiKey)
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}

	// Use custom RoundTripper to fix encoding and parameter issues
	configuration.HTTPClient = &http.Client{
		Transport: &encodingRoundTripper{
			Proxied: http.DefaultTransport,
		},
	}

	return api.NewAPIClient(configuration), context.Background(), viper.GetBool("verbose")
}

func init() {
	// Subcommands will be added in their respective files' init() functions
	// but we can also do it here if we want to be explicit.
}

func handleResponse(cmd *cobra.Command, r *http.Response, err error, res interface{}, isVerbose bool, method string) error {
	if isVerbose && r != nil {
		cmd.Printf("Request URL: %s %s\n", r.Request.Method, r.Request.URL.String())
	}

	if err != nil {
		if isVerbose && r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			return fmt.Errorf("error when calling %s: %w\nFull HTTP response: %v", method, err, r)
		}
		return fmt.Errorf("error when calling %s: %w", method, err)
	}

	jsonRes, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}

	if isVerbose {
		if r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
		}
		cmd.Printf("Response from %s:\n%s\n", method, string(jsonRes))
	} else {
		cmd.Println(string(jsonRes))
	}
	return nil
}
