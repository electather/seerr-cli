package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var requestsCmd = &cobra.Command{
	Use:   "requests <userId>",
	Short: "Get requests for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		take, _ := cmd.Flags().GetFloat32("take")
		skip, _ := cmd.Flags().GetFloat32("skip")

		req := apiClient.UsersAPI.UserUserIdRequestsGet(ctx, float32(userId))
		if cmd.Flags().Changed("take") {
			req = req.Take(take)
		}
		if cmd.Flags().Changed("skip") {
			req = req.Skip(skip)
		}

		res, r, err := req.Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdRequestsGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdRequestsGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdRequestsGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	requestsCmd.Flags().Float32("take", 0, "Number of items to take")
	requestsCmd.Flags().Float32("skip", 0, "Number of items to skip")
	Cmd.AddCommand(requestsCmd)
}
