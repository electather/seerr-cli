package tv

import (
	"strconv"

	"github.com/spf13/cobra"
)

var similarCmd = &cobra.Command{
	Use:   "similar <tvId>",
	Short: "Get similar TV shows",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get similar shows to Breaking Bad (ID 1396)
  seerr-cli tv similar 1396`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		tvId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.TvAPI.TvTvIdSimilarGet(ctx, float32(tvId))
		if cmd.Flags().Changed("page") {
			page, _ := cmd.Flags().GetFloat32("page")
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "TvTvIdSimilarGet")
	},
}

func init() {
	similarCmd.Flags().Float32("page", 1, "Page number")
	similarCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(similarCmd)
}
