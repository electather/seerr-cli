package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	api "seer-cli/pkg/api"
)

var statusAppdataCmd = &cobra.Command{
	Use:   "appdata",
	Short: "Get application data volume status",
	Long:  `For Docker installs, returns whether or not the volume mount was configured properly. Always returns true for non-Docker installs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the API configuration
		configuration := api.NewConfiguration()
		
		serverURL := viper.GetString("server")
		// Ensure server URL has /api/v1 suffix if not present
		if !strings.HasSuffix(serverURL, "/api/v1") && !strings.HasSuffix(serverURL, "/api/v1/") {
			serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
		}

		configuration.Servers = api.ServerConfigurations{
			{
				URL: serverURL,
				Description: "Configured Server",
			},
		}

		if apiKey := viper.GetString("api_key"); apiKey != "" {
			configuration.AddDefaultHeader("X-Api-Key", apiKey)
		}
		
		if overrideServerURL != "" {
			configuration.Servers = api.ServerConfigurations{
				{
					URL: overrideServerURL,
					Description: "Mock Server",
				},
			}
		}
		
		apiClient := api.NewAPIClient(configuration)
		ctx := context.Background()
		
		isVerbose := viper.GetBool("verbose")
		if isVerbose {
			cmd.Printf("Calling /status/appdata endpoint on %s...\n", serverURL)
		}
		
		res, r, err := apiClient.PublicAPI.StatusAppdataGet(ctx).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling StatusAppdataGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling StatusAppdataGet: %w", err)
		}
		
		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from StatusAppdataGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	statusCmd.AddCommand(statusAppdataCmd)
}
