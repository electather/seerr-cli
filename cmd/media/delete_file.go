package media

import (
	"github.com/spf13/cobra"
)

var deleteFileCmd = &cobra.Command{
	Use:   "delete-file <mediaId>",
	Short: "Delete a media file from Radarr/Sonarr",
	Args:  cobra.ExactArgs(1),
	Example: `  # Delete the regular (non-4K) file
  seerr-cli media delete-file 42

  # Delete the 4K file
  seerr-cli media delete-file 42 --is4k`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		req := apiClient.MediaAPI.MediaMediaIdFileDelete(ctx, args[0])
		if cmd.Flags().Changed("is4k") {
			v, _ := cmd.Flags().GetBool("is4k")
			req = req.Is4k(v)
		}

		r, err := req.Execute()
		return handle204Response(cmd, r, err, isVerbose, "MediaMediaIdFileDelete")
	},
}

func init() {
	deleteFileCmd.Flags().Bool("is4k", false, "Remove from 4K service instance instead of regular")
	Cmd.AddCommand(deleteFileCmd)
}
