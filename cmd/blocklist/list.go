package blocklist

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List blocklisted items",
	Example: `  # List all blocklisted items
  seerr-cli blocklist list

  # Filter and paginate
  seerr-cli blocklist list --filter all --take 25 --skip 0`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		req := apiClient.BlocklistAPI.BlocklistGet(ctx)

		if cmd.Flags().Changed("take") {
			v, _ := cmd.Flags().GetInt("take")
			req = req.Take(float32(v))
		}
		if cmd.Flags().Changed("skip") {
			v, _ := cmd.Flags().GetInt("skip")
			req = req.Skip(float32(v))
		}
		if cmd.Flags().Changed("search") {
			v, _ := cmd.Flags().GetString("search")
			req = req.Search(v)
		}
		if cmd.Flags().Changed("filter") {
			v, _ := cmd.Flags().GetString("filter")
			req = req.Filter(v)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "BlocklistGet")
	},
}

func init() {
	listCmd.Flags().Int("take", 20, "Number of items to return")
	listCmd.Flags().Int("skip", 0, "Number of items to skip")
	listCmd.Flags().String("search", "", "Search by title")
	listCmd.Flags().String("filter", "", "Filter: all, manual, blocklistedTags")
	Cmd.AddCommand(listCmd)
}
