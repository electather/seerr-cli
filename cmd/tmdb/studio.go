package tmdb

import (
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var studioCmd = &cobra.Command{
	Use:   "studio <studioId>",
	Short: "Get movie studio details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid studio ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, apiErr := apiClient.TmdbAPI.StudioStudioIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, apiErr, res, isVerbose, "StudioStudioIdGet")
	},
}

func init() {
	Cmd.AddCommand(studioCmd)
}
