package other

import (
	"github.com/spf13/cobra"
)

var watchproviderRegionsCmd = &cobra.Command{
	Use:     "watchprovider-regions",
	Short:   "List all available watch provider regions",
	Example: `  seer-cli other watchprovider-regions`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.OtherAPI.WatchprovidersRegionsGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "WatchprovidersRegionsGet")
	},
}

func init() {
	Cmd.AddCommand(watchproviderRegionsCmd)
}
