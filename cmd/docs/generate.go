package docs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate CLI reference documentation as Markdown",
	Example: `  # Generate docs into the default directory
  seerr-cli docs generate

  # Generate docs into a custom directory
  seerr-cli docs generate --output-dir /tmp/cli-docs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, _ := cmd.Flags().GetString("output-dir")
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory %s: %w", dir, err)
		}
		root := cmd.Root()
		if err := doc.GenMarkdownTree(root, dir); err != nil {
			return fmt.Errorf("failed to generate docs: %w", err)
		}
		cmd.Printf("Documentation written to %s\n", dir)
		return nil
	},
}

func init() {
	generateCmd.Flags().String("output-dir", "./docs/cli/", "Directory to write generated Markdown files")
	Cmd.AddCommand(generateCmd)
}
