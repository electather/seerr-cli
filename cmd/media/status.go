package media

import (
	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <mediaId> <status>",
	Short: "Update media status",
	Args:  cobra.ExactArgs(2),
	Example: `  # Mark media as available
  seerr-cli media status 42 available

  # Mark 4K version as available
  seerr-cli media status 42 available --is4k`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		body := api.NewMediaMediaIdStatusPostRequest()
		if cmd.Flags().Changed("is4k") {
			v, _ := cmd.Flags().GetBool("is4k")
			body.SetIs4k(v)
		}

		res, r, err := apiClient.MediaAPI.MediaMediaIdStatusPost(ctx, args[0], args[1]).
			MediaMediaIdStatusPostRequest(*body).
			Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "MediaMediaIdStatusPost")
	},
}

func init() {
	statusCmd.Flags().Bool("is4k", false, "Update 4K status field instead of regular status")
	Cmd.AddCommand(statusCmd)
}
