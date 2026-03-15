package blocklist

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

var Cmd = &cobra.Command{
	Use:   "blocklist",
	Short: "Manage the blocklist",
	Long:  `Manage blocklisted media items: list, add, get, and remove entries.`,
}

func newAPIClient() (*api.APIClient, context.Context, bool) {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("seerr.server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	if apiKey := viper.GetString("seerr.api_key"); apiKey != "" {
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

func handleRawResponse(cmd *cobra.Command, r *http.Response, err error, isVerbose bool, method string) error {
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
	if isVerbose && r != nil {
		cmd.Printf("HTTP Status: %s\n", r.Status)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer r.Body.Close()
	cmd.Println(string(body))
	return nil
}

func handle204Response(cmd *cobra.Command, r *http.Response, err error, verbose bool, method string) error {
	if err != nil {
		if verbose && r != nil {
			return fmt.Errorf("error when calling %s: %w\nFull HTTP response: %v", method, err, r)
		}
		return fmt.Errorf("error when calling %s: %w", method, err)
	}
	if verbose {
		cmd.Printf("HTTP Status: %s\n", r.Status)
	}
	cmd.Println(`{"status": "ok"}`)
	return nil
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
