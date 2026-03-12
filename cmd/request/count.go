package request

import (
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Get request counts by status",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.RequestAPI.RequestCountGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestCountGet")
	},
}

func init() {
	Cmd.AddCommand(countCmd)
}
