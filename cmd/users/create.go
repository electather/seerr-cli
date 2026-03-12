package users

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		email, _ := cmd.Flags().GetString("email")
		username, _ := cmd.Flags().GetString("username")
		permissions, _ := cmd.Flags().GetFloat32("permissions")

		body := api.UserPostRequest{
			Email: &email,
		}
		if cmd.Flags().Changed("username") {
			body.Username = &username
		}
		if cmd.Flags().Changed("permissions") {
			body.Permissions = &permissions
		}

		res, r, err := apiClient.UsersAPI.UserPost(ctx).UserPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	createCmd.Flags().String("email", "", "Email address (required)")
	createCmd.MarkFlagRequired("email")
	createCmd.Flags().String("username", "", "Username")
	createCmd.Flags().Float32("permissions", 0, "Initial permissions bitmask")
	Cmd.AddCommand(createCmd)
}
