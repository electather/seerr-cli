package search

import (
	"seerr-cli/cmd/apiutil"

	"github.com/spf13/cobra"
)

var tvCmd = &cobra.Command{
	Use:   "tv",
	Short: "Discover and filter TV shows",
	Example: `  # Discover Drama TV shows (Genre ID 18)
  seerr-cli search tv --genre 18

  # Discover TV shows on Netflix (Network ID 213)
  seerr-cli search tv --network 213`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")
		genre, _ := cmd.Flags().GetString("genre")
		network, _ := cmd.Flags().GetInt("network")
		keywords, _ := cmd.Flags().GetString("keywords")
		excludeKeywords, _ := cmd.Flags().GetString("exclude-keywords")
		sortBy, _ := cmd.Flags().GetString("sort-by")
		firstAirGte, _ := cmd.Flags().GetString("first-air-date-gte")
		firstAirLte, _ := cmd.Flags().GetString("first-air-date-lte")

		req := apiClient.SearchAPI.DiscoverTvGet(ctx)
		if cmd.Flags().Changed("page") {
			req = req.Page(float32(page))
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}
		if cmd.Flags().Changed("genre") {
			req = req.Genre(genre)
		}
		if cmd.Flags().Changed("network") {
			req = req.Network(float32(network))
		}
		if cmd.Flags().Changed("keywords") {
			req = req.Keywords(keywords)
		}
		if cmd.Flags().Changed("exclude-keywords") {
			req = req.ExcludeKeywords(excludeKeywords)
		}
		if cmd.Flags().Changed("sort-by") {
			req = req.SortBy(sortBy)
		}
		if cmd.Flags().Changed("first-air-date-gte") {
			req = req.FirstAirDateGte(firstAirGte)
		}
		if cmd.Flags().Changed("first-air-date-lte") {
			req = req.FirstAirDateLte(firstAirLte)
		}

		res, r, err := req.Execute()
		return apiutil.HandleResponse(cmd, r, err, res, isVerbose, "DiscoverTvGet")
	},
}

func init() {
	tvCmd.Flags().Int("page", 1, "Page number")
	tvCmd.Flags().String("language", "en", "Language code")
	tvCmd.Flags().String("genre", "", "Genre ID")
	tvCmd.Flags().Int("network", 0, "Network ID")
	tvCmd.Flags().String("keywords", "", "Keyword IDs")
	tvCmd.Flags().String("exclude-keywords", "", "Keyword IDs to exclude")
	tvCmd.Flags().String("sort-by", "popularity.desc", "Sort by field")
	tvCmd.Flags().String("first-air-date-gte", "", "First air date GTE (YYYY-MM-DD)")
	tvCmd.Flags().String("first-air-date-lte", "", "First air date LTE (YYYY-MM-DD)")
	Cmd.AddCommand(tvCmd)
}
