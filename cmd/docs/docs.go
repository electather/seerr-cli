// Package docs provides commands for generating CLI reference documentation.
package docs

import (
	"github.com/spf13/cobra"
)

// Cmd is the parent command for documentation-related subcommands.
var Cmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate CLI reference documentation",
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
