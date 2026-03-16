package search

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "search",
	Short: "Search for movies, TV shows, people, and more",
	Long:  `Search for various resources from the Seerr API including movies, TV shows, people, keywords, and companies.`,
}

func init() {
	// Subcommands are added in their respective files' init() functions.
}
