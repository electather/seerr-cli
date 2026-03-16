package blocklist

import (
	"seerr-cli/cmd/apiutil"
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add media to the blocklist",
	Example: `  seerr-cli blocklist add --tmdb-id 12345`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		body := api.NewBlocklist()
		if cmd.Flags().Changed("tmdb-id") {
			v, _ := cmd.Flags().GetInt("tmdb-id")
			body.SetTmdbId(float32(v))
		}

		r, err := apiClient.BlocklistAPI.BlocklistPost(ctx).Blocklist(*body).Execute()
		return apiutil.Handle204Response(cmd, r, err, isVerbose, "BlocklistPost")
	},
}

func init() {
	addCmd.Flags().Int("tmdb-id", 0, "TMDB ID of the media to blocklist (required)")
	addCmd.MarkFlagRequired("tmdb-id")
	Cmd.AddCommand(addCmd)
}
