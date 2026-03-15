package other

import (
	"github.com/spf13/cobra"
)

var watchprovidersMoviesCmd = &cobra.Command{
	Use:     "watchproviders-movies",
	Short:   "List watch providers for movies in a region",
	Example: `  seerr-cli other watchproviders-movies --watch-region US`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		req := apiClient.OtherAPI.WatchprovidersMoviesGet(ctx)
		if cmd.Flags().Changed("watch-region") {
			region, _ := cmd.Flags().GetString("watch-region")
			req = req.WatchRegion(region)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "WatchprovidersMoviesGet")
	},
}

func init() {
	watchprovidersMoviesCmd.Flags().String("watch-region", "", "Region code (e.g. US, GB)")
	watchprovidersMoviesCmd.MarkFlagRequired("watch-region")
	Cmd.AddCommand(watchprovidersMoviesCmd)
}
