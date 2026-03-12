package media

import (
	"github.com/spf13/cobra"
)

var watchDataCmd = &cobra.Command{
	Use:   "watch-data <mediaId>",
	Short: "Get watch data for a media item",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get watch data for media item 42
  seer-cli media watch-data 42`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.MediaAPI.MediaMediaIdWatchDataGet(ctx, args[0]).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MediaMediaIdWatchDataGet")
	},
}

func init() {
	Cmd.AddCommand(watchDataCmd)
}
