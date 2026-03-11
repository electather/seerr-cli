package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Persist configuration to the config file",
	Long:  `Save the server URL and API key provided as flags to the CLI configuration file (~/.seer-cli.yaml by default).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// If config file doesn't exist, create it
		if viper.ConfigFileUsed() == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			configPath := filepath.Join(home, ".seer-cli.yaml")
			
			// Set the config file path in viper
			viper.SetConfigFile(configPath)
			
			// Initialize empty file if it doesn't exist
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				if err := viper.SafeWriteConfigAs(configPath); err != nil {
					return fmt.Errorf("failed to create config file: %w", err)
				}
			}
		}

		// Set values from flags explicitly to ensure they are saved
		if s, _ := cmd.Flags().GetString("server"); s != "" {
			viper.Set("server", s)
		} else if s, _ := cmd.Root().PersistentFlags().GetString("server"); s != "" {
			viper.Set("server", s)
		}
		
		if k, _ := cmd.Flags().GetString("api-key"); k != "" {
			viper.Set("api_key", k)
		} else if k, _ := cmd.Root().PersistentFlags().GetString("api-key"); k != "" {
			viper.Set("api_key", k)
		}

		// Save the current viper settings (including flags) to the file
		if err := viper.WriteConfig(); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		cmd.Printf("Configuration saved successfully to: %s\n", viper.ConfigFileUsed())
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(configSetCmd)
}
