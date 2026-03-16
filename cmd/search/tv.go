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

var tvCmd = &cobra.Command{
	Use:   "tv",
	Short: "Discover and filter TV shows",
	Example: `  # Discover Drama TV shows (Genre ID 18)
  seerr-cli search tv --genre 18

  # Discover TV shows on Netflix (Network ID 213)
  seerr-cli search tv --network 213`,
	RunE: func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")
		genre, _ := cmd.Flags().GetString("genre")
		network, _ := cmd.Flags().GetInt("network")
		keywords, _ := cmd.Flags().GetString("keywords")
		excludeKeywords, _ := cmd.Flags().GetString("exclude-keywords")
		sortBy, _ := cmd.Flags().GetString("sort-by")
		firstAirGte, _ := cmd.Flags().GetString("first-air-date-gte")
		firstAirLte, _ := cmd.Flags().GetString("first-air-date-lte")

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
		if cmd.Flags().Changed("network") {
			params.Set("network", strconv.Itoa(network))
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
		if cmd.Flags().Changed("first-air-date-gte") {
			params.Set("firstAirDateGte", firstAirGte)
		}
		if cmd.Flags().Changed("first-air-date-lte") {
			params.Set("firstAirDateLte", firstAirLte)
		}

		b, err := seerrclient.New().DiscoverTV(params)
		if err != nil {
			return err
		}

		if viper.GetBool("verbose") {
			cmd.Printf("GET /api/v1/discover/tv\n")
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
