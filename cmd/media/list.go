package media

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List media items",
	Example: `  # List all media
  seerr-cli media list

  # List available media, sorted by date added
  seerr-cli media list --filter available --sort added

  # Paginate results
  seerr-cli media list --take 10 --skip 20`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		req := apiClient.MediaAPI.MediaGet(ctx)

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

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MediaGet")
	},
}

func init() {
	listCmd.Flags().Int("take", 20, "Number of items to return")
	listCmd.Flags().Int("skip", 0, "Number of items to skip")
	listCmd.Flags().String("filter", "", "Filter by status: all, available, partial, allavailable, processing, pending, deleted")
	listCmd.Flags().String("sort", "", "Sort by: added, modified, mediaAdded")
	Cmd.AddCommand(listCmd)
}
