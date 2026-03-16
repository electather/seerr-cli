package overriderule

import (
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <ruleId>",
	Short: "Update an override rule",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid rule ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, apiErr := apiClient.OverrideruleAPI.OverrideRuleRuleIdPut(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, apiErr, res, isVerbose, "OverrideRuleRuleIdPut")
	},
}

func init() {
	Cmd.AddCommand(updateCmd)
}
