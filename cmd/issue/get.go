package issue

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"seerr-cli/cmd/apiutil"
)

var getCmd = &cobra.Command{
	Use:   "get <issueId>",
	Short: "Get a single issue",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid issue ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, apiErr := apiClient.IssueAPI.IssueIssueIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, apiErr, res, isVerbose, "IssueIssueIdGet")
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
