package cmd

import (
	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check Seer service status",
	Long:  `Check the status of various Seer service components like the system or application data volumes.`,
}

// OverrideServerURL is used by tests to mock the API server
var OverrideServerURL string

func init() {
	RootCmd.AddCommand(StatusCmd)
}
