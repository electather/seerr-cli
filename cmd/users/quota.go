package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var quotaCmd = &cobra.Command{
	Use:   "quota <userId>",
	Short: "Get quota for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdQuotaGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdQuotaGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdQuotaGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdQuotaGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(quotaCmd)
}
