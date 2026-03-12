package other

import (
	"strconv"

	"github.com/spf13/cobra"
)

var keywordCmd = &cobra.Command{
	Use:     "keyword <keywordId>",
	Short:   "Get a keyword by ID",
	Args:    cobra.ExactArgs(1),
	Example: `  seer-cli other keyword 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		keywordId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		res, r, err := apiClient.OtherAPI.KeywordKeywordIdGet(ctx, float32(keywordId)).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "KeywordKeywordIdGet")
	},
}

func init() {
	Cmd.AddCommand(keywordCmd)
}
