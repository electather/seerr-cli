package users

import (
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <userId>",
	Short: "Delete a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		_, r, err := apiClient.UsersAPI.UserUserIdDelete(ctx, float32(userId)).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "UserUserIdDelete")
	},
}

func init() {
	Cmd.AddCommand(deleteCmd)
}
