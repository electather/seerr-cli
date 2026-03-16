package tmdb

import (
	"fmt"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var networkCmd = &cobra.Command{
	Use:   "network <networkId>",
	Short: "Get TV network details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid network ID: %s", args[0])
		}
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, apiErr := apiClient.TmdbAPI.NetworkNetworkIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, apiErr, res, isVerbose, "NetworkNetworkIdGet")
	},
}

func init() {
	Cmd.AddCommand(networkCmd)
}
