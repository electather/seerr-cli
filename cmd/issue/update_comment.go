package issue

import (
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var updateCommentCmd = &cobra.Command{
	Use:   "update-comment <commentId>",
	Short: "Update an issue comment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		body := api.NewIssueCommentCommentIdPutRequest()
		if cmd.Flags().Changed("message") {
			v, _ := cmd.Flags().GetString("message")
			body.SetMessage(v)
		}
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.IssueAPI.IssueCommentCommentIdPut(ctx, args[0]).IssueCommentCommentIdPutRequest(*body).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "IssueCommentCommentIdPut")
	},
}

func init() {
	updateCommentCmd.Flags().String("message", "", "Updated comment message")
	Cmd.AddCommand(updateCommentCmd)
}
