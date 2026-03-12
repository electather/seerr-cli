package service

import (
	"strconv"

	"github.com/spf13/cobra"
)

var radarrListCmd = &cobra.Command{
	Use:   "radarr-list",
	Short: "List Radarr servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.ServiceAPI.ServiceRadarrGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "ServiceRadarrGet")
	},
}

var radarrGetCmd = &cobra.Command{
	Use:   "radarr-get <radarrId>",
	Short: "Get Radarr server quality profiles and root folders",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		id, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}

		res, r, err := apiClient.ServiceAPI.ServiceRadarrRadarrIdGet(ctx, float32(id)).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "ServiceRadarrRadarrIdGet")
	},
}

func init() {
	Cmd.AddCommand(radarrListCmd)
	Cmd.AddCommand(radarrGetCmd)
}
