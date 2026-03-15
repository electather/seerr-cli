package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"seerr-cli/cmd/blocklist"
	"seerr-cli/cmd/collection"
	"seerr-cli/cmd/config"
	"seerr-cli/cmd/issue"
	"seerr-cli/cmd/mcp"
	"seerr-cli/cmd/media"
	"seerr-cli/cmd/movies"
	"seerr-cli/cmd/other"
	"seerr-cli/cmd/overriderule"
	"seerr-cli/cmd/person"
	"seerr-cli/cmd/request"
	"seerr-cli/cmd/search"
	"seerr-cli/cmd/service"
	"seerr-cli/cmd/status"
	"seerr-cli/cmd/tmdb"
	"seerr-cli/cmd/tv"
	"seerr-cli/cmd/users"
	"seerr-cli/cmd/watchlist"
)

var (
	cfgFile string
	server  string
	apiKey  string
	verbose bool

	buildVersion = "dev"
	buildCommit  = "none"
	buildDate    = "unknown"
)

// SetVersionInfo is called from main to inject linker-set build variables.
func SetVersionInfo(version, commit, date string) {
	buildVersion = version
	buildCommit = commit
	buildDate = date

	RootCmd.Version = fmt.Sprintf("%s (commit: %s, built: %s)", buildVersion, buildCommit, buildDate)
}

var RootCmd = &cobra.Command{
	Use:   "seerr-cli",
	Short: "A CLI to interact with the Seerr API",
	Long:  `A command line interface to call endpoints defined in the Seerr OpenAPI specification.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Root() != cmd && (cmd.Name() == "config" || cmd.Parent().Name() == "config" || cmd.Name() == "help" || cmd.Name() == "completion" || cmd.Parent().Name() == "completion" || cmd.Name() == "mcp" || cmd.Parent().Name() == "mcp") {
			return nil
		}
		if viper.GetString("seerr.server") == "" {
			return fmt.Errorf("server URL is required. Set it via --server flag, SEER_SERVER env var, or in the config file")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seerr-cli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "Seerr server URL (e.g., https://seerr.example.com). The /api/v1 prefix is added automatically if not provided.")
	RootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Seerr API Key")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	viper.BindPFlag("seerr.server", RootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("seerr.api_key", RootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

	RootCmd.AddCommand(config.Cmd)
	RootCmd.AddCommand(mcp.Cmd)
	RootCmd.AddCommand(status.Cmd)
	RootCmd.AddCommand(users.Cmd)
	RootCmd.AddCommand(search.Cmd)
	RootCmd.AddCommand(movies.Cmd)
	RootCmd.AddCommand(tv.Cmd)
	RootCmd.AddCommand(other.Cmd)
	RootCmd.AddCommand(person.Cmd)
	RootCmd.AddCommand(request.Cmd)
	RootCmd.AddCommand(media.Cmd)
	RootCmd.AddCommand(blocklist.Cmd)
	RootCmd.AddCommand(watchlist.Cmd)
	RootCmd.AddCommand(collection.Cmd)
	RootCmd.AddCommand(service.Cmd)
	RootCmd.AddCommand(tmdb.Cmd)
	RootCmd.AddCommand(issue.Cmd)
	RootCmd.AddCommand(overriderule.Cmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".seerr-cli")
	}

	viper.SetEnvPrefix("SEER")
	// Replace dots and hyphens so nested keys like "mcp.transport" map to
	// env vars like "SEER_MCP_TRANSPORT".
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
	// AutomaticEnv with the SEERR prefix would construct "SEER_SEER_SERVER" for
	// the "seerr.server" key, so we bind those explicitly instead.
	viper.BindEnv("seerr.server", "SEER_SERVER")
	viper.BindEnv("seerr.api_key", "SEER_API_KEY")

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
