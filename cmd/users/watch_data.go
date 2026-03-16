package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var watchDataCmd = &cobra.Command{
	Use:   "watch-data <userId>",
	Short: "Get watch data for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		res, r, err := apiClient.UsersAPI.UserUserIdWatchDataGet(ctx, float32(userId)).Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdWatchDataGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdWatchDataGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdWatchDataGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(watchDataCmd)
}
