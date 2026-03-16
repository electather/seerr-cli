package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var ratingsCmd = &cobra.Command{
	Use:   "ratings <movieId>",
	Short: "Get movie ratings",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get ratings for The Matrix (ID 603)
  seerr-cli movies ratings 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.MoviesAPI.MovieMovieIdRatingsGet(ctx, float32(movieId)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdRatingsGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCmd)
}
