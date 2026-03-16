package users

import (
	"encoding/json"
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var watchlistCmd = &cobra.Command{
	Use:   "watchlist <userId>",
	Short: "Get watchlist for a user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		userId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid userId: %w", err)
		}

		page, _ := cmd.Flags().GetInt("page")

		req := apiClient.UsersAPI.UserUserIdWatchlistGet(ctx, float32(userId))
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}

		res, r, err := req.Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserUserIdWatchlistGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserUserIdWatchlistGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserUserIdWatchlistGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	watchlistCmd.Flags().Int("page", 1, "Page number to retrieve")
	Cmd.AddCommand(watchlistCmd)
}
