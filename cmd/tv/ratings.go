package tv

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var ratingsCmd = &cobra.Command{
	Use:   "ratings <tvId>",
	Short: "Get TV show ratings",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get ratings for Breaking Bad (ID 1396)
  seerr-cli tv ratings 1396`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		tvId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.TvAPI.TvTvIdRatingsGet(ctx, float32(tvId)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "TvTvIdRatingsGet")
	},
}

func init() {
	Cmd.AddCommand(ratingsCmd)
}
