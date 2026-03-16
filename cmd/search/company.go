package search

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Search for companies",
	Example: `  # Search for "Warner Bros."
  seerr-cli search company -q "Warner Bros."`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetInt("page")

		req := apiClient.SearchAPI.SearchCompanyGet(ctx).Query(query)
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "SearchCompanyGet")
	},
}

func init() {
	companyCmd.Flags().StringP("query", "q", "", "Search query")
	companyCmd.MarkFlagRequired("query")
	companyCmd.Flags().Int("page", 1, "Page number")
	Cmd.AddCommand(companyCmd)
}
