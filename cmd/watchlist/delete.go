package watchlist

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <tmdbId>",
	Short: "Remove an item from the watchlist",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		r, err := apiClient.WatchlistAPI.WatchlistTmdbIdDelete(ctx, args[0]).Execute()
		return handle204Response(cmd, r, err, isVerbose, "WatchlistTmdbIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
