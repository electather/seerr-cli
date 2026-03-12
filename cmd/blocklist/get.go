package blocklist

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <tmdbId>",
	Short: "Get a blocklisted item by TMDB ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		r, err := apiClient.BlocklistAPI.BlocklistTmdbIdGet(ctx, args[0]).Execute()
		return handleRawResponse(cmd, r, err, isVerbose, "BlocklistTmdbIdGet")
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
