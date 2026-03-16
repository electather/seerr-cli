package config

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
		cmd.Printf("Server:      %s\n", viper.GetString("seerr.server"))

		apiKey := viper.GetString("seerr.api_key")
		if apiKey != "" {
			// Show only the last 4 characters so the key is identifiable without
			// exposing the prefix, which is the more sensitive portion.
			var masked string
			if len(apiKey) > 4 {
				masked = "********" + apiKey[len(apiKey)-4:]
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
	Cmd.AddCommand(configShowCmd)
}
