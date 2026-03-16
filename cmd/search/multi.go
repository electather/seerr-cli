package search

import (
	"seerr-cli/internal/seerrclient"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var multiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Search for movies, TV shows, and people",
	Example: `  # Search for "The Matrix"
  seerr-cli search multi -q "The Matrix"

  # Search for "Christopher Nolan" on the second page
  seerr-cli search multi -q "Christopher Nolan" --page 2`,
	RunE: func(cmd *cobra.Command, args []string) error {
		query, _ := cmd.Flags().GetString("query")
		page, _ := cmd.Flags().GetInt("page")
		language, _ := cmd.Flags().GetString("language")

		// Zero out defaults so SearchMulti only includes explicitly-set params.
		if !cmd.Flags().Changed("page") {
			page = 0
		}
		if !cmd.Flags().Changed("language") {
			language = ""
		}

		b, err := seerrclient.New().SearchMulti(query, page, language)
		if err != nil {
			return err
		}

		if viper.GetBool("verbose") {
			cmd.Printf("GET /api/v1/search\n")
		}
		cmd.Println(string(b))
		return nil
	},
}

func init() {
	multiCmd.Flags().StringP("query", "q", "", "Search query")
	multiCmd.MarkFlagRequired("query")
	multiCmd.Flags().Int("page", 1, "Page number")
	multiCmd.Flags().String("language", "en", "Language code")
	Cmd.AddCommand(multiCmd)
}
