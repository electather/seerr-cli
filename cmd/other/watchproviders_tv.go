package other

import (
	"github.com/spf13/cobra"
)

var watchprovidersTvCmd = &cobra.Command{
	Use:     "watchproviders-tv",
	Short:   "List watch providers for TV shows in a region",
	Example: `  seer-cli other watchproviders-tv --watch-region US`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		req := apiClient.OtherAPI.WatchprovidersTvGet(ctx)
		if cmd.Flags().Changed("watch-region") {
			region, _ := cmd.Flags().GetString("watch-region")
			req = req.WatchRegion(region)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "WatchprovidersTvGet")
	},
}

func init() {
	watchprovidersTvCmd.Flags().String("watch-region", "", "Region code (e.g. US, GB)")
	watchprovidersTvCmd.MarkFlagRequired("watch-region")
	Cmd.AddCommand(watchprovidersTvCmd)
}
