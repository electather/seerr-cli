package request

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <requestId>",
	Short: "Delete a media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		r, err := apiClient.RequestAPI.RequestRequestIdDelete(ctx, args[0]).Execute()
		return handle204Response(cmd, r, err, isVerbose, "RequestRequestIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
