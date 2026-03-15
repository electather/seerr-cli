package tv

import (
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <tvId>",
	Short: "Get TV show details",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get details for Breaking Bad (ID 1396)
  seerr-cli tv get 1396

  # Get details in Spanish
  seerr-cli tv get 1396 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		tvId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.TvAPI.TvTvIdGet(ctx, float32(tvId))
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "TvTvIdGet")
	},
}

func init() {
	getCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(getCmd)
}
