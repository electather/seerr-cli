package users

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var importFromPlexCmd = &cobra.Command{
	Use:   "import-from-plex",
	Short: "Import users from Plex",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		plexIds, _ := cmd.Flags().GetStringSlice("plex-ids")

		body := api.UserImportFromPlexPostRequest{
			PlexIds: plexIds,
		}

		res, r, err := apiClient.UsersAPI.UserImportFromPlexPost(ctx).UserImportFromPlexPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserImportFromPlexPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserImportFromPlexPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserImportFromPlexPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var importFromJellyfinCmd = &cobra.Command{
	Use:   "import-from-jellyfin",
	Short: "Import users from Jellyfin",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		jellyfinUserIds, _ := cmd.Flags().GetStringSlice("jellyfin-user-ids")

		body := api.UserImportFromJellyfinPostRequest{
			JellyfinUserIds: jellyfinUserIds,
		}

		res, r, err := apiClient.UsersAPI.UserImportFromJellyfinPost(ctx).UserImportFromJellyfinPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserImportFromJellyfinPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserImportFromJellyfinPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserImportFromJellyfinPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	importFromPlexCmd.Flags().StringSlice("plex-ids", []string{}, "Plex user IDs to import")
	importFromPlexCmd.MarkFlagRequired("plex-ids")

	importFromJellyfinCmd.Flags().StringSlice("jellyfin-user-ids", []string{}, "Jellyfin user IDs to import")
	importFromJellyfinCmd.MarkFlagRequired("jellyfin-user-ids")

	Cmd.AddCommand(importFromPlexCmd)
	Cmd.AddCommand(importFromJellyfinCmd)
}
