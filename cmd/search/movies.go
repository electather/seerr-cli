package search

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")
		genre, _ := cmd.Flags().GetString("genre")
		studio, _ := cmd.Flags().GetInt("studio")
		keywords, _ := cmd.Flags().GetString("keywords")
		excludeKeywords, _ := cmd.Flags().GetString("exclude-keywords")
		sortBy, _ := cmd.Flags().GetString("sort-by")
		releaseGte, _ := cmd.Flags().GetString("release-date-gte")
		releaseLte, _ := cmd.Flags().GetString("release-date-lte")

		params := url.Values{}
		if cmd.Flags().Changed("page") {
			params.Set("page", strconv.Itoa(page))
		}
		if cmd.Flags().Changed("language") {
			params.Set("language", language)
		}
		if cmd.Flags().Changed("genre") {
			params.Set("genre", genre)
		}
		if cmd.Flags().Changed("studio") {
			params.Set("studio", strconv.Itoa(studio))
		}
		if cmd.Flags().Changed("keywords") {
			params.Set("keywords", keywords)
		}
		if cmd.Flags().Changed("exclude-keywords") {
			params.Set("excludeKeywords", excludeKeywords)
		}
		if cmd.Flags().Changed("sort-by") {
			params.Set("sortBy", sortBy)
		}
		if cmd.Flags().Changed("release-date-gte") {
			params.Set("primaryReleaseDateGte", releaseGte)
		}
		if cmd.Flags().Changed("release-date-lte") {
			params.Set("primaryReleaseDateLte", releaseLte)
		}

		b, err := seerrclient.New().DiscoverMovies(params)
		if err != nil {
			return err
		}

		if viper.GetBool("verbose") {
			cmd.Printf("GET /api/v1/discover/movies\n")
		}

		var out interface{}
		if err := json.Unmarshal(b, &out); err != nil {
			return fmt.Errorf("failed to parse response: %w", err)
		}
		formatted, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to format response: %w", err)
		}
		cmd.Println(string(formatted))
		return nil
	},
}

func init() {
	moviesCmd.Flags().Int("page", 1, "Page number")
	moviesCmd.Flags().String("language", "en", "Language code")
	moviesCmd.Flags().String("genre", "", "Genre ID")
	moviesCmd.Flags().Int("studio", 0, "Studio ID")
	moviesCmd.Flags().String("keywords", "", "Keyword IDs")
	moviesCmd.Flags().String("exclude-keywords", "", "Keyword IDs to exclude")
	moviesCmd.Flags().String("sort-by", "popularity.desc", "Sort by field")
	moviesCmd.Flags().String("release-date-gte", "", "Release date GTE (YYYY-MM-DD)")
	moviesCmd.Flags().String("release-date-lte", "", "Release date LTE (YYYY-MM-DD)")
	Cmd.AddCommand(moviesCmd)
}
