package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var similarCmd = &cobra.Command{
	Use:   "similar <movieId>",
	Short: "Get similar movies",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get similar movies to The Matrix (ID 603)
  seerr-cli movies similar 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdSimilarGet(ctx, float32(movieId))
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdSimilarGet")
	},
}

func init() {
	similarCmd.Flags().Int("page", 1, "Page number")
	similarCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(similarCmd)
}
