package request

import (
	"github.com/spf13/cobra"
)

var approveCmd = &cobra.Command{
	Use:   "approve <requestId>",
	Short: "Approve a media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.RequestAPI.RequestRequestIdStatusPost(ctx, args[0], "approve").Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestRequestIdStatusPost")
	},
}

func init() {
	Cmd.AddCommand(approveCmd)
}
