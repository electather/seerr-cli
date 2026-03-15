package search

import (
	"github.com/spf13/cobra"
)

var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Get trending movies and TV shows",
	Example: `  # Get currently trending movies and TV shows
  seerr-cli search trending

  # Get trending items for the week
  seerr-cli search trending --time-window week`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")
		mediaType, _ := cmd.Flags().GetString("media-type")
		timeWindow, _ := cmd.Flags().GetString("time-window")

		req := apiClient.SearchAPI.DiscoverTrendingGet(ctx)
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}
		if cmd.Flags().Changed("media-type") {
			req = req.MediaType(mediaType)
		}
		if cmd.Flags().Changed("time-window") {
			req = req.TimeWindow(timeWindow)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "DiscoverTrendingGet")
	},
}

func init() {
	trendingCmd.Flags().Float32("page", 1, "Page number")
	trendingCmd.Flags().String("language", "en", "Language code")
	trendingCmd.Flags().String("media-type", "all", "Media type (all, movie, tv)")
	trendingCmd.Flags().String("time-window", "day", "Time window (day, week)")
	Cmd.AddCommand(trendingCmd)
}
