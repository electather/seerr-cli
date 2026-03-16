package service

import (
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var sonarrListCmd = &cobra.Command{
	Use:   "sonarr-list",
	Short: "List Sonarr servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()
		res, r, err := apiClient.ServiceAPI.ServiceSonarrGet(ctx).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "ServiceSonarrGet")
	},
}

var sonarrGetCmd = &cobra.Command{
	Use:   "sonarr-get <sonarrId>",
	Short: "Get Sonarr server quality profiles and root folders",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.ServiceAPI.ServiceSonarrSonarrIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "ServiceSonarrSonarrIdGet")
	},
}

var sonarrLookupCmd = &cobra.Command{
	Use:   "sonarr-lookup <tmdbId>",
	Short: "Look up a series in Sonarr by TMDB ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}

		res, r, err := apiClient.ServiceAPI.ServiceSonarrLookupTmdbIdGet(ctx, float32(id)).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "ServiceSonarrLookupTmdbIdGet")
	},
}

func init() {
	Cmd.AddCommand(sonarrListCmd)
	Cmd.AddCommand(sonarrGetCmd)
	Cmd.AddCommand(sonarrLookupCmd)
}
