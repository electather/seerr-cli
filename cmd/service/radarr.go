package service

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var radarrListCmd = &cobra.Command{
	Use:   "radarr-list",
	Short: "List Radarr servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.ServiceAPI.ServiceRadarrGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "ServiceRadarrGet")
	},
}

var radarrGetCmd = &cobra.Command{
	Use:   "radarr-get <radarrId>",
	Short: "Get Radarr server quality profiles and root folders",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.ServiceAPI.ServiceRadarrRadarrIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "ServiceRadarrRadarrIdGet")
	},
}

func init() {
	Cmd.AddCommand(radarrListCmd)
	Cmd.AddCommand(radarrGetCmd)
}
