package tv

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var seasonCmd = &cobra.Command{
	Use:   "season <tvId> <seasonNumber>",
	Short: "Get season details and episode list",
	Args:  cobra.ExactArgs(2),
	Example: `  # Get season 1 of Breaking Bad (ID 1396)
  seerr-cli tv season 1396 1

  # Get season 2 in Spanish
  seerr-cli tv season 1396 2 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		tvId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}
		seasonNumber, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return err
		}

		req := apiClient.TvAPI.TvTvIdSeasonSeasonNumberGet(ctx, float32(tvId), float32(seasonNumber))
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "TvTvIdSeasonSeasonNumberGet")
	},
}

func init() {
	seasonCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(seasonCmd)
}
