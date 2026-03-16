package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ratingsCombinedCmd = &cobra.Command{
	Use:   "ratings-combined <movieId>",
	Short: "Get combined movie ratings (RT/IMDB)",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get combined ratings for The Matrix (ID 603)
  seerr-cli movies ratings-combined 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := seerrclient.New().MovieRatingsCombined(int(movieId))
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "MovieMovieIdRatingscombinedGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCombinedCmd)
}
