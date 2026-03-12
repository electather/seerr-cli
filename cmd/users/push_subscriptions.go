package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var pushSubscriptionsCmd = &cobra.Command{
	Use:   "push-subscriptions",
	Short: "Manage web push subscriptions",
}

var pushSubscriptionsRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a web push subscription",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		endpoint, _ := cmd.Flags().GetString("endpoint")
		auth, _ := cmd.Flags().GetString("auth")
		p256dh, _ := cmd.Flags().GetString("p256dh")
		userAgent, _ := cmd.Flags().GetString("user-agent")

		body := api.UserRegisterPushSubscriptionPostRequest{
			Endpoint: endpoint,
			Auth:     auth,
			P256dh:   p256dh,
		}
		if cmd.Flags().Changed("user-agent") {
			body.UserAgent = &userAgent
		}

		r, err := apiClient.UsersAPI.UserRegisterPushSubscriptionPost(ctx).UserRegisterPushSubscriptionPostRequest(body).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserRegisterPushSubscriptionPost")
	},
}

var pushSubscriptionsListCmd = &cobra.Command{
	Use:   "list <userId>",
	Short: "List push subscriptions for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdPushSubscriptionsGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdPushSubscriptionsGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdPushSubscriptionsGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdPushSubscriptionsGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var pushSubscriptionsGetCmd = &cobra.Command{
	Use:   "get <userId> <endpoint>",
	Short: "Get a specific push subscription for a user",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}
		endpoint := args[1]

		res, r, err := apiClient.UsersAPI.UserUserIdPushSubscriptionEndpointGet(ctx, float32(userId), endpoint).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdPushSubscriptionEndpointGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdPushSubscriptionEndpointGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdPushSubscriptionEndpointGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

var pushSubscriptionsDeleteCmd = &cobra.Command{
	Use:   "delete <userId> <endpoint>",
	Short: "Delete a push subscription for a user",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}
		endpoint := args[1]

		r, err := apiClient.UsersAPI.UserUserIdPushSubscriptionEndpointDelete(ctx, float32(userId), endpoint).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdPushSubscriptionEndpointDelete")
	},
}

func init() {
	pushSubscriptionsRegisterCmd.Flags().String("endpoint", "", "Push endpoint (required)")
	pushSubscriptionsRegisterCmd.MarkFlagRequired("endpoint")
	pushSubscriptionsRegisterCmd.Flags().String("auth", "", "Auth key (required)")
	pushSubscriptionsRegisterCmd.MarkFlagRequired("auth")
	pushSubscriptionsRegisterCmd.Flags().String("p256dh", "", "P256dh key (required)")
	pushSubscriptionsRegisterCmd.MarkFlagRequired("p256dh")
	pushSubscriptionsRegisterCmd.Flags().String("user-agent", "", "User agent")

	pushSubscriptionsCmd.AddCommand(pushSubscriptionsRegisterCmd, pushSubscriptionsListCmd, pushSubscriptionsGetCmd, pushSubscriptionsDeleteCmd)
}
