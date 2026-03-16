package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <userId>",
	Short: "Update an existing user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		permissions, _ := cmd.Flags().GetInt("permissions")

		body := api.UserUpdatePayload{}
		changed := false
		if cmd.Flags().Changed("username") {
			body.Username = &username
			changed = true
		}
		if cmd.Flags().Changed("email") {
			body.Email = &email
			changed = true
		}
		if cmd.Flags().Changed("permissions") {
			p := float32(permissions)
			body.Permissions = &p
			changed = true
		}
		if !changed {
			return fmt.Errorf("at least one field must be provided")
		}

		res, r, err := apiClient.UsersAPI.UserUserIdPut(ctx, float32(userId)).UserUpdatePayload(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdPut: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdPut: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdPut:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	updateCmd.Flags().String("username", "", "New username")
	updateCmd.Flags().String("email", "", "New email address")
	updateCmd.Flags().Int("permissions", 0, "New permissions bitmask")
	Cmd.AddCommand(updateCmd)
}
