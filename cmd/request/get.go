package request

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <requestId>",
	Short: "Get a specific media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.RequestAPI.RequestRequestIdGet(ctx, args[0]).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestRequestIdGet")
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
