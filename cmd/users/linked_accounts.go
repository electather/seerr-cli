package users

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var linkedAccountsCmd = &cobra.Command{
	Use:   "linked-accounts",
	Short: "Manage linked accounts for a user",
}

var linkPlexCmd = &cobra.Command{
	Use:   "link-plex <userId>",
	Short: "Link a Plex account to a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		plexToken, _ := cmd.Flags().GetString("plex-token")
		body := api.AuthPlexPostRequest{
			AuthToken: plexToken,
		}

		r, err := apiClient.UsersAPI.UserUserIdSettingsLinkedAccountsPlexPost(ctx, float32(userId)).AuthPlexPostRequest(body).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdSettingsLinkedAccountsPlexPost")
	},
}

var unlinkPlexCmd = &cobra.Command{
	Use:   "unlink-plex <userId>",
	Short: "Unlink a Plex account from a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		r, err := apiClient.UsersAPI.UserUserIdSettingsLinkedAccountsPlexDelete(ctx, float32(userId)).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdSettingsLinkedAccountsPlexDelete")
	},
}

var linkJellyfinCmd = &cobra.Command{
	Use:   "link-jellyfin <userId>",
	Short: "Link a Jellyfin account to a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		body := api.UserUserIdSettingsLinkedAccountsJellyfinPostRequest{
			Username: &username,
		}
		if cmd.Flags().Changed("password") {
			body.Password = &password
		}

		r, err := apiClient.UsersAPI.UserUserIdSettingsLinkedAccountsJellyfinPost(ctx, float32(userId)).UserUserIdSettingsLinkedAccountsJellyfinPostRequest(body).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdSettingsLinkedAccountsJellyfinPost")
	},
}

var unlinkJellyfinCmd = &cobra.Command{
	Use:   "unlink-jellyfin <userId>",
	Short: "Unlink a Jellyfin account from a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		r, err := apiClient.UsersAPI.UserUserIdSettingsLinkedAccountsJellyfinDelete(ctx, float32(userId)).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdSettingsLinkedAccountsJellyfinDelete")
	},
}

func init() {
	linkPlexCmd.Flags().String("plex-token", "", "Plex authentication token (required)")
	linkPlexCmd.MarkFlagRequired("plex-token")

	linkJellyfinCmd.Flags().String("username", "", "Jellyfin username (required)")
	linkJellyfinCmd.MarkFlagRequired("username")
	linkJellyfinCmd.Flags().String("password", "", "Jellyfin password")

	linkedAccountsCmd.AddCommand(linkPlexCmd, unlinkPlexCmd, linkJellyfinCmd, unlinkJellyfinCmd)
}
