package person

import (
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <personId>",
	Short: "Get person details",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get details for Brad Pitt (ID 287)
  seerr-cli person get 287

  # Get details in Spanish
  seerr-cli person get 287 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		personId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.PersonAPI.PersonPersonIdGet(ctx, float32(personId))
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "PersonPersonIdGet")
	},
}

func init() {
	getCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(getCmd)
}
