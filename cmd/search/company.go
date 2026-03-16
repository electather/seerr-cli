package search

import (
	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Search for companies",
	Example: `  # Search for "Warner Bros."
  seerr-cli search company -q "Warner Bros."`,
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetInt("page")

		sc := seerrclient.New()
		req := sc.Unwrap().SearchAPI.SearchCompanyGet(sc.Ctx()).Query(query)
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "SearchCompanyGet")
	},
}

func init() {
	companyCmd.Flags().StringP("query", "q", "", "Search query")
	companyCmd.MarkFlagRequired("query")
	companyCmd.Flags().Int("page", 1, "Page number")
	Cmd.AddCommand(companyCmd)
}
