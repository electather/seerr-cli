package users

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var bulkUpdateCmd = &cobra.Command{
	Use:   "bulk-update",
	Short: "Update permissions for multiple users",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		ids, _ := cmd.Flags().GetInt32Slice("ids")
		permissions, _ := cmd.Flags().GetInt32("permissions")

		body := api.UserPutRequest{
			Ids: ids,
		}
		if cmd.Flags().Changed("permissions") {
			body.Permissions = &permissions
		}

		res, r, err := apiClient.UsersAPI.UserPut(ctx).UserPutRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserPut: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserPut: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserPut:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	bulkUpdateCmd.Flags().Int32Slice("ids", []int32{}, "User IDs to update (required)")
	bulkUpdateCmd.MarkFlagRequired("ids")
	bulkUpdateCmd.Flags().Int32("permissions", 0, "New permissions bitmask (required)")
	bulkUpdateCmd.MarkFlagRequired("permissions")
	Cmd.AddCommand(bulkUpdateCmd)
}
