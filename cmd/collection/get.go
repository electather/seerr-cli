package collection

import (
	"strconv"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <collectionId>",
	Short: "Get collection details",
	Args:  cobra.ExactArgs(1),
	Example: `  seerr-cli collection get 537982
  seerr-cli collection get 537982 --language en`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		id, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		req := apiClient.CollectionAPI.CollectionCollectionIdGet(ctx, float32(id))
		if cmd.Flags().Changed("language") {
			v, _ := cmd.Flags().GetString("language")
			req = req.Language(v)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "CollectionCollectionIdGet")
	},
}

func init() {
	getCmd.Flags().String("language", "", "Language code (e.g. en)")
	Cmd.AddCommand(getCmd)
}
