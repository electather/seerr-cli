package search

import (
	"net/url"
	"strconv"

	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Search for movies, TV shows, and people",
	Example: `  # Search for "The Matrix"
  seerr-cli search multi -q "The Matrix"

  # Search for "Christopher Nolan" on the second page
  seerr-cli search multi -q "Christopher Nolan" --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")

		params := url.Values{}
		params.Set("query", query)
		if cmd.Flags().Changed("page") {
			params.Set("page", strconv.Itoa(int(page)))
		}
		if cmd.Flags().Changed("language") {
			params.Set("language", language)
		}

		// Use a raw HTTP request to avoid the broken union-type unmarshal in the
		// generated client (TV results are incorrectly parsed as PersonResult) and
		// to ensure spaces are encoded as %20 rather than + in the query string.
		b, err := apiutil.RawGet(ctx, apiClient, "/search", params)
		if err != nil {
			return err
		}

		if isVerbose {
			cmd.Printf("GET /api/v1/search\n")
		}
		cmd.Println(string(b))
		return nil
	},
}

func init() {
	multiCmd.Flags().StringP("query", "q", "", "Search query")
	multiCmd.MarkFlagRequired("query")
	multiCmd.Flags().Float32("page", 1, "Page number")
	multiCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(multiCmd)
}
