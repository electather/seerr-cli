package request

import (
	"fmt"
	"strconv"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <requestId>",
	Short: "Update a media request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		mediaType, _ := cmd.Flags().GetString("media-type")
		body := *api.NewRequestRequestIdPutRequest(mediaType)

		if cmd.Flags().Changed("seasons") {
			v, _ := cmd.Flags().GetString("seasons")
			parts := strings.Split(v, ",")
			nums := make([]float32, 0, len(parts))
			for _, p := range parts {
				n, err := strconv.ParseFloat(strings.TrimSpace(p), 32)
				if err != nil {
					return fmt.Errorf("invalid season number %q: %w", p, err)
				}
				nums = append(nums, float32(n))
			}
			body.SetSeasons(nums)
		}
		if cmd.Flags().Changed("is4k") {
			v, _ := cmd.Flags().GetBool("is4k")
			body.SetIs4k(v)
		}
		if cmd.Flags().Changed("server-id") {
			v, _ := cmd.Flags().GetFloat32("server-id")
			body.SetServerId(v)
		}
		if cmd.Flags().Changed("profile-id") {
			v, _ := cmd.Flags().GetFloat32("profile-id")
			body.SetProfileId(v)
		}
		if cmd.Flags().Changed("root-folder") {
			v, _ := cmd.Flags().GetString("root-folder")
			body.SetRootFolder(v)
		}
		if cmd.Flags().Changed("language-profile-id") {
			v, _ := cmd.Flags().GetFloat32("language-profile-id")
			body.SetLanguageProfileId(v)
		}
		if cmd.Flags().Changed("user-id") {
			v, _ := cmd.Flags().GetFloat32("user-id")
			body.SetUserId(v)
		}

		res, r, err := apiClient.RequestAPI.RequestRequestIdPut(ctx, args[0]).RequestRequestIdPutRequest(body).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestRequestIdPut")
	},
}

func init() {
	updateCmd.Flags().String("media-type", "", "Media type: movie or tv (required)")
	updateCmd.MarkFlagRequired("media-type")
	updateCmd.Flags().String("seasons", "", `Comma-separated season numbers (e.g. "1,2,3")`)
	updateCmd.Flags().Bool("is4k", false, "Request 4K version")
	updateCmd.Flags().Float32("server-id", 0, "Target server ID")
	updateCmd.Flags().Float32("profile-id", 0, "Quality profile ID")
	updateCmd.Flags().String("root-folder", "", "Root folder path")
	updateCmd.Flags().Float32("language-profile-id", 0, "Language profile ID")
	updateCmd.Flags().Float32("user-id", 0, "User ID")
	Cmd.AddCommand(updateCmd)
}
