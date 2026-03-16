package issue

import (
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var commentCmd = &cobra.Command{
	Use:   "comment <issueId>",
	Short: "Add a comment to an issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid issue ID: %s", args[0])
		}
		message, _ := cmd.Flags().GetString("message")
		body := api.NewIssueIssueIdCommentPostRequest(message)
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, apiErr := apiClient.IssueAPI.IssueIssueIdCommentPost(ctx, float32(id)).IssueIssueIdCommentPostRequest(*body).Execute()
		return apiutil.HandleResponse(cmd, r, apiErr, res, isVerbose, "IssueIssueIdCommentPost")
	},
}

func init() {
	commentCmd.Flags().String("message", "", "Comment message")
	commentCmd.MarkFlagRequired("message")
	Cmd.AddCommand(commentCmd)
}
