package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Manage user passwords",
}

var passwordResetRequestCmd = &cobra.Command{
	Use:   "reset-request",
	Short: "Request a password reset",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		email, _ := cmd.Flags().GetString("email")

		body := api.AuthResetPasswordPostRequest{
			Email: email,
		}

		res, r, err := apiClient.UsersAPI.AuthResetPasswordPost(ctx).AuthResetPasswordPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling AuthResetPasswordPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling AuthResetPasswordPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from AuthResetPasswordPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var passwordResetConfirmCmd = &cobra.Command{
	Use:   "reset-confirm <guid>",
	Short: "Confirm a password reset",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		guid := args[0]
		password, _ := cmd.Flags().GetString("password")

		body := api.AuthResetPasswordGuidPostRequest{
			Password: password,
		}

		res, r, err := apiClient.UsersAPI.AuthResetPasswordGuidPost(ctx, guid).AuthResetPasswordGuidPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling AuthResetPasswordGuidPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling AuthResetPasswordGuidPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from AuthResetPasswordGuidPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var passwordGetCmd = &cobra.Command{
	Use:   "get <userId>",
	Short: "Check if a user has a password set",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsPasswordGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsPasswordGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsPasswordGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsPasswordGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var passwordSetCmd = &cobra.Command{
	Use:   "set <userId>",
	Short: "Set a user's password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		newPassword, _ := cmd.Flags().GetString("new-password")
		currentPassword, _ := cmd.Flags().GetString("current-password")

		body := api.UserUserIdSettingsPasswordPostRequest{
			NewPassword: newPassword,
		}
		if cmd.Flags().Changed("current-password") {
			body.CurrentPassword = *api.NewNullableString(&currentPassword)
		} else {
			body.CurrentPassword = *api.NewNullableString(nil)
		}

		r, err := apiClient.UsersAPI.UserUserIdSettingsPasswordPost(ctx, float32(userId)).UserUserIdSettingsPasswordPostRequest(body).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "UserUserIdSettingsPasswordPost")
	},
}

func init() {
	passwordResetRequestCmd.Flags().String("email", "", "Email address for reset (required)")
	passwordResetRequestCmd.MarkFlagRequired("email")

	passwordResetConfirmCmd.Flags().String("password", "", "New password (required)")
	passwordResetConfirmCmd.MarkFlagRequired("password")

	passwordSetCmd.Flags().String("new-password", "", "New password (required)")
	passwordSetCmd.MarkFlagRequired("new-password")
	passwordSetCmd.Flags().String("current-password", "", "Current password (if required)")

	passwordCmd.AddCommand(passwordResetRequestCmd, passwordResetConfirmCmd, passwordGetCmd, passwordSetCmd)
}
