package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		res, r, err := seerrclient.New().MovieRecommendations(int(movieId), page, language)
		return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "MovieMovieIdRecommendationsGet")
	},
}

func init() {
	recommendationsCmd.Flags().Int("page", 1, "Page number")
	recommendationsCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(recommendationsCmd)
}
