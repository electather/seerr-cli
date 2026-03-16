package other

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var keywordCmd = &cobra.Command{
	Use:     "keyword <keywordId>",
	Short:   "Get a keyword by ID",
	Args:    cobra.ExactArgs(1),
	Example: `  seerr-cli other keyword 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		keywordId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.OtherAPI.KeywordKeywordIdGet(ctx, float32(keywordId)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "KeywordKeywordIdGet")
	},
}

func init() {
	Cmd.AddCommand(keywordCmd)
}
