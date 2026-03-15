package request

import (
	"fmt"
	"strconv"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new media request",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		mediaType, _ := cmd.Flags().GetString("media-type")
		mediaId, _ := cmd.Flags().GetFloat32("media-id")
		body := *api.NewRequestPostRequest(mediaType, mediaId)

		if cmd.Flags().Changed("tvdb-id") {
			v, _ := cmd.Flags().GetFloat32("tvdb-id")
			body.SetTvdbId(v)
		}
		if cmd.Flags().Changed("seasons") {
			v, _ := cmd.Flags().GetString("seasons")
			if v == "all" {
				body.SetSeasons(api.StringAsRequestPostRequestSeasons(&v))
			} else {
				parts := strings.Split(v, ",")
				nums := make([]float32, 0, len(parts))
				for _, p := range parts {
					n, err := strconv.ParseFloat(strings.TrimSpace(p), 32)
					if err != nil {
						return fmt.Errorf("invalid season number %q: %w", p, err)
					}
					nums = append(nums, float32(n))
				}
				body.SetSeasons(api.ArrayOfFloat32AsRequestPostRequestSeasons(&nums))
			}
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

		res, r, err := apiClient.RequestAPI.RequestPost(ctx).RequestPostRequest(body).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "RequestPost")
	},
}

func init() {
	createCmd.Flags().String("media-type", "", "Media type: movie or tv (required)")
	createCmd.MarkFlagRequired("media-type")
	createCmd.Flags().Float32("media-id", 0, "The TMDB media ID (required)")
	createCmd.MarkFlagRequired("media-id")
	createCmd.Flags().Float32("tvdb-id", 0, "TVDB ID")
	createCmd.Flags().String("seasons", "", `Seasons to request: "all" or comma-separated numbers (e.g. "1,2,3")`)
	createCmd.Flags().Bool("is4k", false, "Request 4K version")
	createCmd.Flags().Float32("server-id", 0, "Target server ID")
	createCmd.Flags().Float32("profile-id", 0, "Quality profile ID")
	createCmd.Flags().String("root-folder", "", "Root folder path")
	createCmd.Flags().Float32("language-profile-id", 0, "Language profile ID")
	createCmd.Flags().Float32("user-id", 0, "User ID to create request as")
	Cmd.AddCommand(createCmd)
}
