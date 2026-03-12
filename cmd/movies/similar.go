package movies

import (
	"strconv"

	"github.com/spf13/cobra"
)

var similarCmd = &cobra.Command{
	Use:   "similar <movieId>",
	Short: "Get similar movies",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get similar movies to The Matrix (ID 603)
  seer-cli movies similar 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		movieId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdSimilarGet(ctx, float32(movieId))
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdSimilarGet")
	},
}

func init() {
	similarCmd.Flags().Float32("page", 1, "Page number")
	similarCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(similarCmd)
}
