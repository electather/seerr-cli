package movies

import (
	"strconv"

	"seerr-cli/cmd/apiutil"
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get <movieId>",
	Short: "Get movie details",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get details for The Matrix (ID 603)
  seerr-cli movies get 603

  # Get details in Spanish
  seerr-cli movies get 603 --language es

  # Output as YAML
  seerr-cli movies get 603 --output yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		movieId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		language, _ := cmd.Flags().GetString("language")
		if !cmd.Flags().Changed("language") {
			language = ""
		}

		mode := apiutil.GetOutputMode(cmd)
		res, r, err := seerrclient.New().MovieGet(int(movieId), language)
		if err != nil {
			return apiutil.HandleResponse(cmd, r, err, res, viper.GetBool("verbose"), "MovieMovieIdGet")
		}

		if viper.GetBool("verbose") && r != nil {
			cmd.Printf("HTTP Status: %s\n", r.Status)
		}
		return apiutil.PrintOutput(cmd, res, mode)
	},
}

func init() {
	getCmd.Flags().String("language", "en", "Language code")
	apiutil.AddOutputFlag(getCmd)
	Cmd.AddCommand(getCmd)
}
