package watchlist

import (
	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add media to the watchlist",
	Example: `  seer-cli watchlist add --tmdb-id 12345 --media-type movie --title "Dune"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		body := api.NewWatchlist()
		if cmd.Flags().Changed("tmdb-id") {
			v, _ := cmd.Flags().GetFloat32("tmdb-id")
			body.SetTmdbId(v)
		}
		if cmd.Flags().Changed("media-type") {
			v, _ := cmd.Flags().GetString("media-type")
			body.SetType(v)
		}
		if cmd.Flags().Changed("title") {
			v, _ := cmd.Flags().GetString("title")
			body.SetTitle(v)
		}
		if cmd.Flags().Changed("rating-key") {
			v, _ := cmd.Flags().GetString("rating-key")
			body.SetRatingKey(v)
		}

		res, r, err := apiClient.WatchlistAPI.WatchlistPost(ctx).Watchlist(*body).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "WatchlistPost")
	},
}

func init() {
	addCmd.Flags().Float32("tmdb-id", 0, "TMDB ID of the media (required)")
	addCmd.MarkFlagRequired("tmdb-id")
	addCmd.Flags().String("media-type", "", "Media type: movie or tv (required)")
	addCmd.MarkFlagRequired("media-type")
	addCmd.Flags().String("title", "", "Title of the media")
	addCmd.Flags().String("rating-key", "", "Plex rating key")
	Cmd.AddCommand(addCmd)
}
