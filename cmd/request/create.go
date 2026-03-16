package request

import (
	"fmt"
	"seerr-cli/cmd/apiutil"
	"strconv"
	"strings"

	api "seerr-cli/pkg/api"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new media request",
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := apiutil.NewAPIClient()

		mediaType, _ := cmd.Flags().GetString("media-type")
		mediaId, _ := cmd.Flags().GetInt("media-id")
		body := *api.NewRequestPostRequest(mediaType, float32(mediaId))

		if cmd.Flags().Changed("tvdb-id") {
			v, _ := cmd.Flags().GetInt("tvdb-id")
			body.SetTvdbId(float32(v))
		}
		if cmd.Flags().Changed("seasons") {
			v, _ := cmd.Flags().GetString("seasons")
			if v == "all" {
				body.SetSeasons(api.StringAsRequestPostRequestSeasons(&v))
			} else {
				parts := strings.Split(v, ",")
				nums := make([]float32, 0, len(parts))
				for _, p := range parts {
					n, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64)
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
			v, _ := cmd.Flags().GetInt("server-id")
			body.SetServerId(float32(v))
		}
		if cmd.Flags().Changed("profile-id") {
			v, _ := cmd.Flags().GetInt("profile-id")
			body.SetProfileId(float32(v))
		}
		if cmd.Flags().Changed("root-folder") {
			v, _ := cmd.Flags().GetString("root-folder")
			body.SetRootFolder(v)
		}
		if cmd.Flags().Changed("language-profile-id") {
			v, _ := cmd.Flags().GetInt("language-profile-id")
			body.SetLanguageProfileId(float32(v))
		}
		if cmd.Flags().Changed("user-id") {
			v, _ := cmd.Flags().GetInt("user-id")
			body.SetUserId(float32(v))
		}

		res, r, err := apiClient.RequestAPI.RequestPost(ctx).RequestPostRequest(body).Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "RequestPost")
	},
}

func init() {
	createCmd.Flags().String("media-type", "", "Media type: movie or tv (required)")
	createCmd.MarkFlagRequired("media-type")
	createCmd.Flags().Int("media-id", 0, "The TMDB media ID (required)")
	createCmd.MarkFlagRequired("media-id")
	createCmd.Flags().Int("tvdb-id", 0, "TVDB ID")
	createCmd.Flags().String("seasons", "", `Seasons to request: "all" or comma-separated numbers (e.g. "1,2,3")`)
	createCmd.Flags().Bool("is4k", false, "Request 4K version")
	createCmd.Flags().Int("server-id", 0, "Target server ID")
	createCmd.Flags().Int("profile-id", 0, "Quality profile ID")
	createCmd.Flags().String("root-folder", "", "Root folder path")
	createCmd.Flags().Int("language-profile-id", 0, "Language profile ID")
	createCmd.Flags().Int("user-id", 0, "User ID to create request as")
	Cmd.AddCommand(createCmd)
}
