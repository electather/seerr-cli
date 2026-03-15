package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Manage user settings",
}

var settingsGetCmd = &cobra.Command{
	Use:   "get <userId>",
	Short: "Get settings for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsMainGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsMainGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsMainGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsMainGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var settingsUpdateCmd = &cobra.Command{
	Use:   "update <userId>",
	Short: "Update settings for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		body := api.UserSettings{}

		if cmd.Flags().Changed("username") {
			v, _ := cmd.Flags().GetString("username")
			body.SetUsername(v)
		}
		if cmd.Flags().Changed("email") {
			v, _ := cmd.Flags().GetString("email")
			body.SetEmail(v)
		}
		if cmd.Flags().Changed("discord-id") {
			v, _ := cmd.Flags().GetString("discord-id")
			body.SetDiscordId(v)
		}
		if cmd.Flags().Changed("locale") {
			v, _ := cmd.Flags().GetString("locale")
			body.SetLocale(v)
		}
		if cmd.Flags().Changed("discover-region") {
			v, _ := cmd.Flags().GetString("discover-region")
			body.SetDiscoverRegion(v)
		}
		if cmd.Flags().Changed("streaming-region") {
			v, _ := cmd.Flags().GetString("streaming-region")
			body.SetStreamingRegion(v)
		}
		if cmd.Flags().Changed("original-language") {
			v, _ := cmd.Flags().GetString("original-language")
			body.SetOriginalLanguage(v)
		}
		if cmd.Flags().Changed("movie-quota-limit") {
			v, _ := cmd.Flags().GetFloat32("movie-quota-limit")
			body.SetMovieQuotaLimit(v)
		}
		if cmd.Flags().Changed("movie-quota-days") {
			v, _ := cmd.Flags().GetFloat32("movie-quota-days")
			body.SetMovieQuotaDays(v)
		}
		if cmd.Flags().Changed("tv-quota-limit") {
			v, _ := cmd.Flags().GetFloat32("tv-quota-limit")
			body.SetTvQuotaLimit(v)
		}
		if cmd.Flags().Changed("tv-quota-days") {
			v, _ := cmd.Flags().GetFloat32("tv-quota-days")
			body.SetTvQuotaDays(v)
		}
		if cmd.Flags().Changed("watchlist-sync-movies") {
			v, _ := cmd.Flags().GetBool("watchlist-sync-movies")
			body.SetWatchlistSyncMovies(v)
		}
		if cmd.Flags().Changed("watchlist-sync-tv") {
			v, _ := cmd.Flags().GetBool("watchlist-sync-tv")
			body.SetWatchlistSyncTv(v)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsMainPost(ctx, float32(userId)).UserSettings(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsMainPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsMainPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsMainPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var settingsNotificationsGetCmd = &cobra.Command{
	Use:   "notifications-get <userId>",
	Short: "Get notification settings for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsNotificationsGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsNotificationsGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsNotificationsGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsNotificationsGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var settingsNotificationsUpdateCmd = &cobra.Command{
	Use:   "notifications-update <userId>",
	Short: "Update notification settings for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		jsonStr, _ := cmd.Flags().GetString("json")
		var body api.UserSettingsNotifications
		if err := json.Unmarshal([]byte(jsonStr), &body); err != nil {
			return fmt.Errorf("invalid JSON: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsNotificationsPost(ctx, float32(userId)).UserSettingsNotifications(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsNotificationsPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsNotificationsPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsNotificationsPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var settingsPermissionsGetCmd = &cobra.Command{
	Use:   "permissions-get <userId>",
	Short: "Get permissions for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsPermissionsGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsPermissionsGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsPermissionsGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsPermissionsGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var settingsPermissionsSetCmd = &cobra.Command{
	Use:   "permissions-set <userId>",
	Short: "Set permissions for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		permissions, _ := cmd.Flags().GetFloat32("permissions")
		body := api.UserUserIdSettingsPermissionsPostRequest{
			Permissions: permissions,
		}

		res, r, err := apiClient.UsersAPI.UserUserIdSettingsPermissionsPost(ctx, float32(userId)).UserUserIdSettingsPermissionsPostRequest(body).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdSettingsPermissionsPost: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdSettingsPermissionsPost: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdSettingsPermissionsPost:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	settingsUpdateCmd.Flags().String("username", "", "New username")
	settingsUpdateCmd.Flags().String("email", "", "New email address")
	settingsUpdateCmd.Flags().String("discord-id", "", "New Discord ID")
	settingsUpdateCmd.Flags().String("locale", "", "New locale")
	settingsUpdateCmd.Flags().String("discover-region", "", "New discover region")
	settingsUpdateCmd.Flags().String("streaming-region", "", "New streaming region")
	settingsUpdateCmd.Flags().String("original-language", "", "New original language")
	settingsUpdateCmd.Flags().Float32("movie-quota-limit", 0, "New movie quota limit")
	settingsUpdateCmd.Flags().Float32("movie-quota-days", 0, "New movie quota days")
	settingsUpdateCmd.Flags().Float32("tv-quota-limit", 0, "New TV quota limit")
	settingsUpdateCmd.Flags().Float32("tv-quota-days", 0, "New TV quota days")
	settingsUpdateCmd.Flags().Bool("watchlist-sync-movies", false, "Enable/disable movie watchlist sync")
	settingsUpdateCmd.Flags().Bool("watchlist-sync-tv", false, "Enable/disable TV watchlist sync")

	settingsNotificationsUpdateCmd.Flags().String("json", "", "Notification settings JSON (required)")
	settingsNotificationsUpdateCmd.MarkFlagRequired("json")

	settingsPermissionsSetCmd.Flags().Float32("permissions", 0, "Permissions bitmask (required)")
	settingsPermissionsSetCmd.MarkFlagRequired("permissions")

	settingsCmd.AddCommand(settingsGetCmd, settingsUpdateCmd, settingsNotificationsGetCmd, settingsNotificationsUpdateCmd, settingsPermissionsGetCmd, settingsPermissionsSetCmd)
}
