package media

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List media items",
	Example: `  # List all media
  seer-cli media list

  # List available media, sorted by date added
  seer-cli media list --filter available --sort added

  # Paginate results
  seer-cli media list --take 10 --skip 20`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		req := apiClient.MediaAPI.MediaGet(ctx)

		if cmd.Flags().Changed("take") {
			v, _ := cmd.Flags().GetFloat32("take")
			req = req.Take(v)
		}
		if cmd.Flags().Changed("skip") {
			v, _ := cmd.Flags().GetFloat32("skip")
			req = req.Skip(v)
		}
		if cmd.Flags().Changed("filter") {
			v, _ := cmd.Flags().GetString("filter")
			req = req.Filter(v)
		}
		if cmd.Flags().Changed("sort") {
			v, _ := cmd.Flags().GetString("sort")
			req = req.Sort(v)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MediaGet")
	},
}

func init() {
	listCmd.Flags().Float32("take", 20, "Number of items to return")
	listCmd.Flags().Float32("skip", 0, "Number of items to skip")
	listCmd.Flags().String("filter", "", "Filter by status: all, available, partial, allavailable, processing, pending, deleted")
	listCmd.Flags().String("sort", "", "Sort by: added, modified, mediaAdded")
	Cmd.AddCommand(listCmd)
}
