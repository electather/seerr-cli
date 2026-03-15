package status

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusSystemCmd = &cobra.Command{
	Use:   "system",
	Short: "Get the system status of the Seerr API",
	Long:  `Call the status endpoint defined in the OpenAPI specification to get the service version and commit details.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configuration := api.NewConfiguration()

		serverURL := viper.GetString("seerr.server")
		if !strings.HasSuffix(serverURL, "/api/v1") && !strings.HasSuffix(serverURL, "/api/v1/") {
			serverURL = strings.TrimSuffix(serverURL, "/") + "/api/v1"
		}

		configuration.Servers = api.ServerConfigurations{
			{URL: serverURL, Description: "Configured Server"},
		}

		if apiKey := viper.GetString("seerr.api_key"); apiKey != "" {
			configuration.AddDefaultHeader("X-Api-Key", apiKey)
		}

		if OverrideServerURL != "" {
			configuration.Servers = api.ServerConfigurations{
				{URL: OverrideServerURL, Description: "Mock Server"},
			}
		}

		apiClient := api.NewAPIClient(configuration)
		ctx := context.Background()

		isVerbose := viper.GetBool("verbose")
		if isVerbose {
			cmd.Printf("Calling /status endpoint on %s...\n", serverURL)
		}

		res, r, err := apiClient.PublicAPI.StatusGet(ctx).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling StatusGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling StatusGet: %w", err)
		}

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
	Cmd.AddCommand(statusSystemCmd)
}
