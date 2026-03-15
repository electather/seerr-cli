package collection

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

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

var Cmd = &cobra.Command{
	Use:   "collection",
	Short: "Manage collections",
	Long:  `Access collection details from the Seer API.`,
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
	return api.NewAPIClient(configuration), context.Background(), viper.GetBool("verbose")
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

func init() {
	// Subcommands are added in their respective files' init() functions.
}
