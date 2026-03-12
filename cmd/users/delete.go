package users

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <userId>",
	Short: "Delete a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		userId, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		_, r, err := apiClient.UsersAPI.UserUserIdDelete(ctx, float32(userId)).Execute()
		return handle204Response(cmd, r, err, isVerbose, "UserUserIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
