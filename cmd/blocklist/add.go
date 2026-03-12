package blocklist

import (
	"github.com/spf13/cobra"
	api "seer-cli/pkg/api"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add media to the blocklist",
	Example: `  seer-cli blocklist add --tmdb-id 12345`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		body := api.NewBlocklist()
		if cmd.Flags().Changed("tmdb-id") {
			v, _ := cmd.Flags().GetFloat32("tmdb-id")
			body.SetTmdbId(v)
		}

		r, err := apiClient.BlocklistAPI.BlocklistPost(ctx).Blocklist(*body).Execute()
		return handle204Response(cmd, r, err, isVerbose, "BlocklistPost")
	},
}

func init() {
	addCmd.Flags().Float32("tmdb-id", 0, "TMDB ID of the media to blocklist (required)")
	addCmd.MarkFlagRequired("tmdb-id")
	Cmd.AddCommand(addCmd)
}
