package request

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all media requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		req := apiClient.RequestAPI.RequestGet(ctx)

		if cmd.Flags().Changed("take") {
			v, _ := cmd.Flags().GetInt("take")
			req = req.Take(float32(v))
		}
		if cmd.Flags().Changed("skip") {
			v, _ := cmd.Flags().GetInt("skip")
			req = req.Skip(float32(v))
		}
		if cmd.Flags().Changed("filter") {
			v, _ := cmd.Flags().GetString("filter")
			req = req.Filter(v)
		}
		if cmd.Flags().Changed("sort") {
			v, _ := cmd.Flags().GetString("sort")
			req = req.Sort(v)
		}
		if cmd.Flags().Changed("sort-direction") {
			v, _ := cmd.Flags().GetString("sort-direction")
			req = req.SortDirection(v)
		}
		if cmd.Flags().Changed("requested-by") {
			v, _ := cmd.Flags().GetInt("requested-by")
			req = req.RequestedBy(float32(v))
		}
		if cmd.Flags().Changed("media-type") {
			v, _ := cmd.Flags().GetString("media-type")
			req = req.MediaType(v)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "RequestGet")
	},
}

func init() {
	listCmd.Flags().Int("take", 0, "Number of results to return")
	listCmd.Flags().Int("skip", 0, "Number of results to skip")
	listCmd.Flags().String("filter", "", "Filter by status (all, approved, available, pending, processing, unavailable, failed, deleted, completed)")
	listCmd.Flags().String("sort", "added", "Sort field (added, modified)")
	listCmd.Flags().String("sort-direction", "desc", "Sort direction (asc, desc)")
	listCmd.Flags().Int("requested-by", 0, "Filter by requesting user ID")
	listCmd.Flags().String("media-type", "all", "Filter by media type (movie, tv, all)")
	Cmd.AddCommand(listCmd)
}
