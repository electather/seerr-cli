package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var ratingsCombinedCmd = &cobra.Command{
	Use:   "ratings-combined <movieId>",
	Short: "Get combined movie ratings (RT/IMDB)",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get combined ratings for The Matrix (ID 603)
  seerr-cli movies ratings-combined 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.MoviesAPI.MovieMovieIdRatingscombinedGet(ctx, float32(movieId)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdRatingscombinedGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCombinedCmd)
}
