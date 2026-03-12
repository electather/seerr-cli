package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"seer-cli/cmd/config"
	"seer-cli/cmd/movies"
	"seer-cli/cmd/other"
	"seer-cli/cmd/person"
	"seer-cli/cmd/request"
	"seer-cli/cmd/search"
	"seer-cli/cmd/status"
	"seer-cli/cmd/tv"
	"seer-cli/cmd/users"
)

var (
	cfgFile string
	server  string
	apiKey  string
	verbose bool
)

var RootCmd = &cobra.Command{
	Use:   "seer-cli",
	Short: "A CLI to interact with the Seer API",
	Long:  `A command line interface to call endpoints defined in the Seer OpenAPI specification.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Root() != cmd && (cmd.Name() == "config" || cmd.Parent().Name() == "config" || cmd.Name() == "help" || cmd.Name() == "completion") {
			return nil
		}
		if viper.GetString("server") == "" {
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.seer-cli.yaml)")
	RootCmd.PersistentFlags().StringVarP(&server, "server", "s", "", "Seer server URL (e.g., https://request.omidastaraki.com). The /api/v1 prefix is added automatically if not provided.")
	RootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Seer API Key")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	viper.BindPFlag("server", RootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("api_key", RootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

	RootCmd.AddCommand(config.Cmd)
	RootCmd.AddCommand(status.Cmd)
	RootCmd.AddCommand(users.Cmd)
	RootCmd.AddCommand(search.Cmd)
	RootCmd.AddCommand(movies.Cmd)
	RootCmd.AddCommand(tv.Cmd)
	RootCmd.AddCommand(other.Cmd)
	RootCmd.AddCommand(person.Cmd)
	RootCmd.AddCommand(request.Cmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".seer-cli")
	}

	viper.SetEnvPrefix("SEER")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}
}
