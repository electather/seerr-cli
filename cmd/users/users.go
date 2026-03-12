package users

import (
	"context"
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
	Use:   "users",
	Short: "Manage users and user settings",
	Long:  `Manage users, their permissions, settings, linked accounts, and push subscriptions.`,
}

func newAPIClient() (*api.APIClient, context.Context, bool) {
	configuration := api.NewConfiguration()
	serverURL := viper.GetString("server")
	if !strings.HasSuffix(serverURL, "/api/v1") {
		serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
	}
	configuration.Servers = api.ServerConfigurations{{URL: serverURL, Description: "Configured Server"}}
	if apiKey := viper.GetString("api_key"); apiKey != "" {
		configuration.AddDefaultHeader("X-Api-Key", apiKey)
	}
	if OverrideServerURL != "" {
		configuration.Servers = api.ServerConfigurations{{URL: OverrideServerURL, Description: "Mock Server"}}
	}
	return api.NewAPIClient(configuration), context.Background(), viper.GetBool("verbose")
}

func init() {
	Cmd.AddCommand(passwordCmd, settingsCmd, linkedAccountsCmd, pushSubscriptionsCmd)
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
