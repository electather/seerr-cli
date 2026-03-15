package search

import (
	"github.com/spf13/cobra"
)

var moviesCmd = &cobra.Command{
	Use:   "movies",
	Short: "Discover and filter movies",
	Example: `  # Discover Drama movies (Genre ID 18)
  seerr-cli search movies --genre 18

  # Discover movies from Warner Bros. (Studio ID 174)
  seerr-cli search movies --studio 174

  # Discover movies sorted by release date
  seerr-cli search movies --sort-by primary_release_date.desc`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiClient, ctx, isVerbose := newAPIClient()

		page, _ := cmd.Flags().GetFloat32("page")
		language, _ := cmd.Flags().GetString("language")
		genre, _ := cmd.Flags().GetString("genre")
		studio, _ := cmd.Flags().GetFloat32("studio")
		keywords, _ := cmd.Flags().GetString("keywords")
		excludeKeywords, _ := cmd.Flags().GetString("exclude-keywords")
		sortBy, _ := cmd.Flags().GetString("sort-by")
		releaseGte, _ := cmd.Flags().GetString("release-date-gte")
		releaseLte, _ := cmd.Flags().GetString("release-date-lte")

		req := apiClient.SearchAPI.DiscoverMoviesGet(ctx)
		if cmd.Flags().Changed("page") {
			req = req.Page(page)
		}
		if cmd.Flags().Changed("language") {
			req = req.Language(language)
		}
		if cmd.Flags().Changed("genre") {
			req = req.Genre(genre)
		}
		if cmd.Flags().Changed("studio") {
			req = req.Studio(studio)
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
		if cmd.Flags().Changed("release-date-gte") {
			req = req.PrimaryReleaseDateGte(releaseGte)
		}
		if cmd.Flags().Changed("release-date-lte") {
			req = req.PrimaryReleaseDateLte(releaseLte)
		}

		res, r, err := req.Execute()
		return handleResponse(cmd, r, err, res, isVerbose, "DiscoverMoviesGet")
	},
}

func init() {
	moviesCmd.Flags().Float32("page", 1, "Page number")
	moviesCmd.Flags().String("language", "en", "Language code")
	moviesCmd.Flags().String("genre", "", "Genre ID")
	moviesCmd.Flags().Float32("studio", 0, "Studio ID")
	moviesCmd.Flags().String("keywords", "", "Keyword IDs")
	moviesCmd.Flags().String("exclude-keywords", "", "Keyword IDs to exclude")
	moviesCmd.Flags().String("sort-by", "popularity.desc", "Sort by field")
	moviesCmd.Flags().String("release-date-gte", "", "Release date GTE (YYYY-MM-DD)")
	moviesCmd.Flags().String("release-date-lte", "", "Release date LTE (YYYY-MM-DD)")
	Cmd.AddCommand(moviesCmd)
}
