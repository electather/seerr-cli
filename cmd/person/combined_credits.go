package person

import (
	"strconv"

	"github.com/spf13/cobra"
)

var combinedCreditsCmd = &cobra.Command{
	Use:   "combined-credits <personId>",
	Short: "Get combined movie and TV credits for a person",
	Args:  cobra.ExactArgs(1),
	Example: `  # Get credits for Brad Pitt (ID 287)
  seer-cli person combined-credits 287

  # Get credits in Spanish
  seer-cli person combined-credits 287 --language es`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		personId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.PersonAPI.PersonPersonIdCombinedCreditsGet(ctx, float32(personId))
		if cmd.Flags().Changed("language") {
			language, _ := cmd.Flags().GetString("language")
			req = req.Language(language)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "PersonPersonIdCombinedCreditsGet")
	},
}

func init() {
	combinedCreditsCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(combinedCreditsCmd)
}
