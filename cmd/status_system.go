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

var statusSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Get the system status of the Seer API",
	Long:  `Call the status endpoint defined in the OpenAPI specification to get the service version and commit details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Initialize the API configuration
		configuration := api.NewConfiguration()
		
		serverURL := viper.GetString("server")
		// Ensure server URL has /api/v1 suffix if not present
		if !strings.HasSuffix(serverURL, "/api/v1") && !strings.HasSuffix(serverURL, "/api/v1/") {
			serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
		}

		// Set server URL from viper
		configuration.Servers = api.ServerConfigurations{
			{
				URL: serverURL,
				Description: "Configured Server",
			},
		}

		// Set API Key from viper if provided
		if apiKey := viper.GetString("api_key"); apiKey != "" {
			configuration.AddDefaultHeader("X-Api-Key", apiKey)
		}
		
		// If overridden by tests, use the mock server
		if OverrideServerURL != "" {
			configuration.Servers = api.ServerConfigurations{
				{
					URL: OverrideServerURL,
					Description: "Mock Server",
				},
			}
		}
		
		apiClient := api.NewAPIClient(configuration)
		
		// Example context
		ctx := context.Background()
		
		isVerbose := viper.GetBool("verbose")
		if isVerbose {
			cmd.Printf("Calling /status endpoint on %s...\n", serverURL)
		}
		
		// Call a status or public endpoint if one exists
		res, r, err := apiClient.PublicAPI.StatusGet(ctx).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling StatusGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling StatusGet: %w", err)
		}
		
		// Pretty print the response as JSON
		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from StatusGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	StatusCmd.AddCommand(statusSystemCmd)
}
