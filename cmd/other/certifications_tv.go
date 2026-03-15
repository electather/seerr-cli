package other

import (
	"github.com/spf13/cobra"
)

var certificationsTvCmd = &cobra.Command{
	Use:     "certifications-tv",
	Short:   "List TV certifications from TMDB",
	Example: `  seerr-cli other certifications-tv`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()
		res, r, err := apiClient.OtherAPI.CertificationsTvGet(ctx).Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "CertificationsTvGet")
	},
}

func init() {
	Cmd.AddCommand(certificationsTvCmd)
}
