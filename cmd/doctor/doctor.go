// Package doctor provides the "doctor" command which verifies that the CLI is
// correctly configured and can reach the configured Seerr server.
package doctor

import (
	"github.com/spf13/cobra"
)

// Cmd is the parent command for all doctor subcommands.
var Cmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check CLI configuration and server connectivity",
	Long:  `Run a series of checks to verify that seerr-cli is correctly configured and can reach the Seerr API.`,
	Example: `  # Run all checks (human-readable output)
  seerr-cli doctor

  # Output check results as JSON
  seerr-cli doctor --output json`,
}

func init() {
	// The check subcommand is added in check.go's init().
}
