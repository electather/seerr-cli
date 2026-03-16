package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var similarCmd = &cobra.Command{
	Use:   "similar <movieId>",
	Short: "Get similar movies",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get similar movies to The Matrix (ID 603)
  seerr-cli movies similar 603`,
	RunE: func(cmd *cobra.Command, args []string) error {
		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		page, _ := cmd.Flags().GetInt("page")
		if !cmd.Flags().Changed("page") {
			page = 0
		}
		language, _ := cmd.Flags().GetString("language")
		if !cmd.Flags().Changed("language") {
			language = ""
		}

		res, r, err := seerrclient.New().MovieSimilar(int(movieId), page, language)
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "MovieMovieIdSimilarGet")
	},
}

func init() {
	similarCmd.Flags().Int("page", 1, "Page number")
	similarCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(similarCmd)
}
