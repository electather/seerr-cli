package users

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		take, _ := cmd.Flags().GetFloat32("take")
		skip, _ := cmd.Flags().GetFloat32("skip")
		sort, _ := cmd.Flags().GetString("sort")
		query, _ := cmd.Flags().GetString("q")
		includeIds, _ := cmd.Flags().GetString("include-ids")

		req := apiClient.UsersAPI.UserGet(ctx)
		if cmd.Flags().Changed("take") {
			req = req.Take(take)
		}
		if cmd.Flags().Changed("skip") {
			req = req.Skip(skip)
		}
		if cmd.Flags().Changed("sort") {
			req = req.Sort(sort)
		}
		if cmd.Flags().Changed("q") {
			req = req.Q(query)
		}
		if cmd.Flags().Changed("include-ids") {
			req = req.IncludeIds(includeIds)
		}

		res, r, err := req.Execute()
		if err != nil {
			if isVerbose && r != nil {
				return fmt.Errorf("error when calling UserGet: %w\nFull HTTP response: %v", err, r)
			}
			return fmt.Errorf("error when calling UserGet: %w", err)
		}

		jsonRes, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}

		if isVerbose {
			cmd.Printf("HTTP Status: %s\n", r.Status)
			cmd.Printf("Response from UserGet:\n%s\n", string(jsonRes))
		} else {
			cmd.Println(string(jsonRes))
		}
		return nil
	},
}

func init() {
	listCmd.Flags().Float32("take", 0, "Number of items to take")
	listCmd.Flags().Float32("skip", 0, "Number of items to skip")
	listCmd.Flags().String("sort", "id", "Field to sort by")
	listCmd.Flags().String("q", "", "Search query")
	listCmd.Flags().String("include-ids", "", "Comma-separated IDs to include")
	Cmd.AddCommand(listCmd)
}
