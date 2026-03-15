package movies

import (
	"strconv"

	"github.com/spf13/cobra"
)

var ratingsCmd = &cobra.Command{
	Use:   "ratings <movieId>",
	Short: "Get movie ratings",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get ratings for The Matrix (ID 603)
  seerr-cli movies ratings 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		movieId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		res, r, err := apiClient.MoviesAPI.MovieMovieIdRatingsGet(ctx, float32(movieId)).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdRatingsGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCmd)
}
