package issue

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all issues",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		req := apiClient.IssueAPI.IssueGet(ctx)
		if cmd.Flags().Changed("take") {
			v, _ := cmd.Flags().GetInt("take")
			req = req.Take(float32(v))
		}
		if cmd.Flags().Changed("skip") {
			v, _ := cmd.Flags().GetInt("skip")
			req = req.Skip(float32(v))
		}
		if cmd.Flags().Changed("sort") {
			v, _ := cmd.Flags().GetString("sort")
			req = req.Sort(v)
		}
		if cmd.Flags().Changed("filter") {
			v, _ := cmd.Flags().GetString("filter")
			req = req.Filter(v)
		}
		if cmd.Flags().Changed("requested-by") {
			v, _ := cmd.Flags().GetInt("requested-by")
			req = req.RequestedBy(float32(v))
		}
		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "IssueGet")
	},
}

func init() {
	listCmd.Flags().Int("take", 20, "Number of issues to return")
	listCmd.Flags().Int("skip", 0, "Number of issues to skip")
	listCmd.Flags().String("sort", "added", "Sort by: added, modified")
	listCmd.Flags().String("filter", "open", "Filter by status: all, open, resolved")
	listCmd.Flags().Int("requested-by", 0, "Filter by user ID")
	Cmd.AddCommand(listCmd)
}
