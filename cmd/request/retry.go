package request

import (
	"github.com/spf13/cobra"
)

var retryCmd = &cobra.Command{
	Use:   "retry <requestId>",
	Short: "Retry a failed media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.RequestAPI.RequestRequestIdRetryPost(ctx, args[0]).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestRequestIdRetryPost")
	},
}

func init() {
	Cmd.AddCommand(retryCmd)
}
