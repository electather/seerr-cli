package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ratingsCmd = &cobra.Command{
	Use:   "ratings <movieId>",
	Short: "Get movie ratings",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get ratings for The Matrix (ID 603)
  seerr-cli movies ratings 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := seerrclient.New().MovieRatings(int(movieId))
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "MovieMovieIdRatingsGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCmd)
}
