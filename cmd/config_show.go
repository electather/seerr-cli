package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the current configuration",
	Long:  `Show the currently active server URL and API key from flags, environment variables, or the config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Config File: %s\n", viper.ConfigFileUsed())
		cmd.Printf("Server:      %s\n", viper.GetString("server"))
		
		apiKey := viper.GetString("api_key")
		if apiKey != "" {
			// Mask API key for security
			masked := apiKey
			if len(apiKey) > 4 {
				masked = apiKey[:4] + "****" + apiKey[len(apiKey)-4:]
			} else {
				masked = "****"
			}
			cmd.Printf("API Key:     %s\n", masked)
		} else {
			cmd.Printf("API Key:     <not set>\n")
		}
	},
}

func init() {
	ConfigCmd.AddCommand(configShowCmd)
}
