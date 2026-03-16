package search

import (
	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var keywordCmd = &cobra.Command{
	Use:   "keyword",
	Short: "Search for keywords",
	Example: `  # Search for the "sci-fi" keyword
  seerr-cli search keyword -q "sci-fi"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetInt("page")

		sc := seerrclient.New()
		req := sc.Unwrap().SearchAPI.SearchKeywordGet(sc.Ctx()).Query(query)
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "SearchKeywordGet")
	},
}

func init() {
	keywordCmd.Flags().StringP("query", "q", "", "Search query")
	keywordCmd.MarkFlagRequired("query")
	keywordCmd.Flags().Int("page", 1, "Page number")
	Cmd.AddCommand(keywordCmd)
}
