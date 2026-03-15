package tv

import (
	"strconv"

	"github.com/spf13/cobra"
)

var recommendationsCmd = &cobra.Command{
	Use:   "recommendations <tvId>",
	Short: "Get recommended TV shows",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get recommendations based on Breaking Bad (ID 1396)
  seerr-cli tv recommendations 1396

  # Get second page
  seerr-cli tv recommendations 1396 --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		tvId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.TvAPI.TvTvIdRecommendationsGet(ctx, float32(tvId))
		if cmd.Flags().Changed("page") {
			page, _ := cmd.Flags().GetFloat32("page")
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "TvTvIdRecommendationsGet")
	},
}

func init() {
	recommendationsCmd.Flags().Float32("page", 1, "Page number")
	recommendationsCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(recommendationsCmd)
}
