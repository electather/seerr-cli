package search

import (
	"github.com/spf13/cobra"
)

var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Search for movies, TV shows, and people",
	Example: `  # Search for "The Matrix"
  seerr-cli search multi -q "The Matrix"

  # Search for "Christopher Nolan" on the second page
  seerr-cli search multi -q "Christopher Nolan" --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")

		req := apiClient.SearchAPI.SearchGet(ctx).Query(query)
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "SearchGet")
	},
}

func init() {
	multiCmd.Flags().StringP("query", "q", "", "Search query")
	multiCmd.MarkFlagRequired("query")
	multiCmd.Flags().Float32("page", 1, "Page number")
	multiCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(multiCmd)
}
