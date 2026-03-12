package movies

import (
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <movieId>",
	Short: "Get movie details",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get details for The Matrix (ID 603)
  seer-cli movies get 603

  # Get details in Spanish
  seer-cli movies get 603 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		movieId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdGet(ctx, float32(movieId))
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdGet")
	},
}

func init() {
	getCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(getCmd)
}
