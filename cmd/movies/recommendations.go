package movies

import (
	"strconv"

	"github.com/spf13/cobra"
)

var recommendationsCmd = &cobra.Command{
	Use:   "recommendations <movieId>",
	Short: "Get recommended movies",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get recommendations based on The Matrix (ID 603)
  seer-cli movies recommendations 603

  # Get second page of recommendations
  seer-cli movies recommendations 603 --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		movieId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdRecommendationsGet(ctx, float32(movieId))
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdRecommendationsGet")
	},
}

func init() {
	recommendationsCmd.Flags().Float32("page", 1, "Page number")
	recommendationsCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(recommendationsCmd)
}
