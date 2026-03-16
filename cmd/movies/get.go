package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <movieId>",
	Short: "Get movie details",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get details for The Matrix (ID 603)
  seerr-cli movies get 603

  # Get details in Spanish
  seerr-cli movies get 603 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		language, _ := cmd.Flags().GetString("language")

		req := apiClient.MoviesAPI.MovieMovieIdGet(ctx, float32(movieId))
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "MovieMovieIdGet")
	},
}

func init() {
	getCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(getCmd)
}
