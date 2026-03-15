package search

import (
	"github.com/spf13/cobra"
)

var keywordCmd = &cobra.Command{
	Use:   "keyword",
	Short: "Search for keywords",
	Example: `  # Search for the "sci-fi" keyword
  seerr-cli search keyword -q "sci-fi"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetFloat32("page")

		req := apiClient.SearchAPI.SearchKeywordGet(ctx).Query(query)
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "SearchKeywordGet")
	},
}

func init() {
	keywordCmd.Flags().StringP("query", "q", "", "Search query")
	keywordCmd.MarkFlagRequired("query")
	keywordCmd.Flags().Float32("page", 1, "Page number")
	Cmd.AddCommand(keywordCmd)
}
