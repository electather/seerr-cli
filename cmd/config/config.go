package config

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
	Long:  `View or update the configuration for the Seerr CLI.`,
}
