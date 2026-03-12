package other

import (
	"github.com/spf13/cobra"
)

var certificationsMovieCmd = &cobra.Command{
	Use:     "certifications-movie",
	Short:   "List movie certifications from TMDB",
	Example: `  seer-cli other certifications-movie`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.OtherAPI.CertificationsMovieGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "CertificationsMovieGet")
	},
}

func init() {
	Cmd.AddCommand(certificationsMovieCmd)
}
