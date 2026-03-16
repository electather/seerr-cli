package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var recommendationsCmd = &cobra.Command{
	Use:   "recommendations <movieId>",
	Short: "Get recommended movies",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get recommendations based on The Matrix (ID 603)
  seerr-cli movies recommendations 603

  # Get second page of recommendations
  seerr-cli movies recommendations 603 --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdRecommendationsGet(ctx, float32(movieId))
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdRecommendationsGet")
	},
}

func init() {
	recommendationsCmd.Flags().Int("page", 1, "Page number")
	recommendationsCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(recommendationsCmd)
}
