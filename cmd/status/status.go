package status

import (
	"github.com/spf13/cobra"
)

// OverrideServerURL is used by tests to redirect API calls to a mock server.
var OverrideServerURL string

var Cmd = &cobra.Command{
	Use:   "status",
	Short: "Check Seerr service status",
	Long:  `Check the status of various Seerr service components like the system or application data volumes.`,
}
