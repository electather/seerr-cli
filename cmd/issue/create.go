package issue

import (
	"fmt"
	"strconv"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <issueType>",
	Short: "Create a new issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		issueType, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid issue type: %s", args[0])
		}
		body := api.NewIssuePostRequest()
		body.SetIssueType(float32(issueType))
		if cmd.Flags().Changed("message") {
			v, _ := cmd.Flags().GetString("message")
			body.SetMessage(v)
		}
		if cmd.Flags().Changed("media-id") {
			v, _ := cmd.Flags().GetFloat32("media-id")
			body.SetMediaId(v)
		}
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, apiErr := apiClient.IssueAPI.IssuePost(ctx).IssuePostRequest(*body).Execute()
		return handleResponse(cmd, r, apiErr, res, isVerbose, "IssuePost")
	},
}

func init() {
	createCmd.Flags().String("message", "", "Issue message")
	createCmd.Flags().Float32("media-id", 0, "Media ID associated with the issue")
	Cmd.AddCommand(createCmd)
}
