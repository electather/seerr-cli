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

var trendingCmd = &cobra.Command{
	Use:   "trending",
	Short: "Get trending movies and TV shows",
	Example: `  # Get currently trending movies and TV shows
  seerr-cli search trending

  # Get trending items for the week
  seerr-cli search trending --time-window week`,
	RunE: func(cmd *cobra.Command, args []string) error {
		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")
		mediaType, _ := cmd.Flags().GetString("media-type")
		timeWindow, _ := cmd.Flags().GetString("time-window")

		params := url.Values{}
		if cmd.Flags().Changed("page") {
			params.Set("page", strconv.Itoa(page))
		}
		if cmd.Flags().Changed("language") {
			params.Set("language", language)
		}
		if cmd.Flags().Changed("media-type") {
			params.Set("mediaType", mediaType)
		}
		if cmd.Flags().Changed("time-window") {
			params.Set("timeWindow", timeWindow)
		}

		b, err := seerrclient.New().DiscoverTrending(params)
		if err != nil {
			return err
		}

		if viper.GetBool("verbose") {
			cmd.Printf("GET /api/v1/discover/trending\n")
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
	trendingCmd.Flags().Int("page", 1, "Page number")
	trendingCmd.Flags().String("language", "en", "Language code")
	trendingCmd.Flags().String("media-type", "all", "Media type (all, movie, tv)")
	trendingCmd.Flags().String("time-window", "day", "Time window (day, week)")
	Cmd.AddCommand(trendingCmd)
}
