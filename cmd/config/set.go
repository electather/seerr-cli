package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const schemaComment = "# yaml-language-server: $schema=https://raw.githubusercontent.com/electather/seer-cli/main/seer-cli.schema.json\n"

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Persist configuration to the config file",
	Long:  `Save the server URL and API key provided as flags to the CLI configuration file (~/.seer-cli.yaml by default).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := viper.ConfigFileUsed()
		if configPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			configPath = filepath.Join(home, ".seer-cli.yaml")
			viper.SetConfigFile(configPath)
		}

		if s, _ := cmd.Root().PersistentFlags().GetString("server"); s != "" {
			viper.Set("seer.server", s)
		}

		if k, _ := cmd.Root().PersistentFlags().GetString("api-key"); k != "" {
			viper.Set("seer.api_key", k)
		}

		if err := writeStructuredConfig(configPath); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		cmd.Printf("Configuration saved successfully to: %s\n", configPath)
		return nil
	},
}

// writeStructuredConfig writes only non-empty config values as structured YAML
// with a yaml-language-server schema comment for IDE autocomplete support.
func writeStructuredConfig(path string) error {
	cfg := buildConfig()

	raw, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	content := schemaComment + string(raw)
	return os.WriteFile(path, []byte(content), 0600)
}

// configFile is the serialisation model for the YAML config file.
// Struct field order determines the YAML key order in the written file.
// Fields tagged with omitempty are omitted when they hold their zero value.
type configFile struct {
	Seer *seerSection `yaml:"seer,omitempty"`
	MCP  *mcpSection  `yaml:"mcp,omitempty"`
}

type seerSection struct {
	Server string `yaml:"server,omitempty"`
	APIKey string `yaml:"api_key,omitempty"`
}

type mcpSection struct {
	Transport   string `yaml:"transport,omitempty"`
	Addr        string `yaml:"addr,omitempty"`
	AuthToken   string `yaml:"auth_token,omitempty"`
	NoAuth      bool   `yaml:"no_auth,omitempty"`
	RouteToken  string `yaml:"route_token,omitempty"`
	TLSCert     string `yaml:"tls_cert,omitempty"`
	TLSKey      string `yaml:"tls_key,omitempty"`
	CORS        bool   `yaml:"cors,omitempty"`
	MultiTenant bool   `yaml:"multi_tenant,omitempty"`
	LogFile     string `yaml:"log_file,omitempty"`
	LogLevel    string `yaml:"log_level,omitempty"`
	LogFormat   string `yaml:"log_format,omitempty"`
}

// buildConfig assembles the config from Viper, omitting default and empty
// values so the written file stays minimal.
func buildConfig() configFile {
	var cfg configFile

	seer := &seerSection{}
	if v := viper.GetString("seer.server"); v != "" {
		seer.Server = strings.TrimRight(v, "/")
	}
	if v := viper.GetString("seer.api_key"); v != "" {
		seer.APIKey = v
	}
	if seer.Server != "" || seer.APIKey != "" {
		cfg.Seer = seer
	}

	mcp := &mcpSection{}
	if v := viper.GetString("mcp.transport"); v != "" && v != "stdio" {
		mcp.Transport = v
	}
	if v := viper.GetString("mcp.addr"); v != "" && v != ":8811" {
		mcp.Addr = v
	}
	if v := viper.GetString("mcp.auth_token"); v != "" {
		mcp.AuthToken = v
	}
	mcp.NoAuth = viper.GetBool("mcp.no_auth")
	if v := viper.GetString("mcp.route_token"); v != "" {
		mcp.RouteToken = v
	}
	if v := viper.GetString("mcp.tls_cert"); v != "" {
		mcp.TLSCert = v
	}
	if v := viper.GetString("mcp.tls_key"); v != "" {
		mcp.TLSKey = v
	}
	mcp.CORS = viper.GetBool("mcp.cors")
	mcp.MultiTenant = viper.GetBool("mcp.multi_tenant")
	if v := viper.GetString("mcp.log_file"); v != "" {
		mcp.LogFile = v
	}
	if v := viper.GetString("mcp.log_level"); v != "" && v != "info" {
		mcp.LogLevel = v
	}
	if v := viper.GetString("mcp.log_format"); v != "" && v != "text" {
		mcp.LogFormat = v
	}
	// Only write the mcp section when at least one non-default value is present.
	if *mcp != (mcpSection{}) {
		cfg.MCP = mcp
	}

	return cfg
}

func init() {
	Cmd.AddCommand(configSetCmd)
}
