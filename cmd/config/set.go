package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"seerr-cli/cmd/mcp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

const schemaComment = "# yaml-language-server: $schema=https://raw.githubusercontent.com/electather/seerr-cli/main/seerr-cli.schema.json\n"

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Persist configuration to the config file",
	Long: `Save configuration to the CLI config file (~/.seerr-cli.yaml by default).

Accepts the global --server and --api-key flags for Seerr instance settings,
and all MCP server flags (same as 'mcp serve') for MCP settings.`,
	Example: `  # Set Seerr instance
  seerr-cli config set --server https://seerr.example.com --api-key mykey

  # Set Seerr instance and configure the MCP server for HTTP transport
  seerr-cli config set --server https://seerr.example.com --api-key mykey \
    --transport http --auth-token mysecret`,
	RunE: func(cmd *cobra.Command, args []string) error {
		configPath := viper.ConfigFileUsed()
		if configPath == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			configPath = filepath.Join(home, ".seerr-cli.yaml")
			viper.SetConfigFile(configPath)
		}

		// Promote explicitly-passed root flags into Viper.
		if s, _ := cmd.Root().PersistentFlags().GetString("server"); s != "" {
			viper.Set("seerr.server", s)
		}
		if k, _ := cmd.Root().PersistentFlags().GetString("api-key"); k != "" {
			viper.Set("seerr.api_key", k)
		}

		// Promote explicitly-passed MCP flags into Viper without touching the
		// serve command's Viper bindings. We only propagate flags that the user
		// actually provided (flag.Changed), leaving existing config-file values
		// for everything else.
		for _, f := range mcp.ServeFlags {
			flag := cmd.Flags().Lookup(f.Name)
			if flag == nil || !flag.Changed {
				continue
			}
			if f.IsBool {
				val, _ := cmd.Flags().GetBool(f.Name)
				viper.Set(f.ViperKey, val)
			} else {
				val, _ := cmd.Flags().GetString(f.Name)
				viper.Set(f.ViperKey, val)
			}
		}

		if err := writeStructuredConfig(configPath); err != nil {
			return fmt.Errorf("failed to save configuration: %w", err)
		}

		cmd.Printf("Configuration saved successfully to: %s\n", configPath)
		return nil
	},
}

// writeStructuredConfig writes only non-empty/non-default config values as
// structured YAML with a yaml-language-server schema comment.
func writeStructuredConfig(path string) error {
	root := &yaml.Node{Kind: yaml.MappingNode}

	if seerNode := buildSeerNode(); seerNode != nil {
		root.Content = append(root.Content,
			scalarNode("seer"),
			seerNode,
		)
	}

	if mcpNode := buildMCPNode(); mcpNode != nil {
		root.Content = append(root.Content,
			scalarNode("mcp"),
			mcpNode,
		)
	}

	if len(root.Content) == 0 {
		// Nothing to write — produce an empty document.
		return os.WriteFile(path, []byte(schemaComment), 0600)
	}

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(&yaml.Node{Kind: yaml.DocumentNode, Content: []*yaml.Node{root}}); err != nil {
		return err
	}
	enc.Close()

	return os.WriteFile(path, []byte(schemaComment+buf.String()), 0600)
}

// buildSeerNode returns a YAML mapping node for the seer: section, or nil if
// neither server nor api_key is set.
func buildSeerNode() *yaml.Node {
	server := strings.TrimRight(viper.GetString("seerr.server"), "/")
	apiKey := viper.GetString("seerr.api_key")

	if server == "" && apiKey == "" {
		return nil
	}

	node := &yaml.Node{Kind: yaml.MappingNode}
	if server != "" {
		node.Content = append(node.Content, scalarNode("server"), scalarNode(server))
	}
	if apiKey != "" {
		node.Content = append(node.Content, scalarNode("api_key"), scalarNode(apiKey))
	}
	return node
}

// buildMCPNode returns a YAML mapping node for the mcp: section driven entirely
// by mcp.ServeFlags. A key is included only when viper.IsSet reports the value
// was explicitly configured (flag, env var, or config file) — defaults are never
// written. Returns nil when no MCP setting is active.
func buildMCPNode() *yaml.Node {
	node := &yaml.Node{Kind: yaml.MappingNode}

	for _, f := range mcp.ServeFlags {
		if !viper.IsSet(f.ViperKey) {
			continue
		}
		// YAML key: take the part of the Viper key after the first dot.
		yamlKey := f.ViperKey
		if idx := strings.Index(f.ViperKey, "."); idx >= 0 {
			yamlKey = f.ViperKey[idx+1:]
		}

		var valNode *yaml.Node
		if f.IsBool {
			valNode = boolNode(viper.GetBool(f.ViperKey))
		} else {
			valNode = scalarNode(viper.GetString(f.ViperKey))
		}
		node.Content = append(node.Content, scalarNode(yamlKey), valNode)
	}

	if len(node.Content) == 0 {
		return nil
	}
	return node
}

// scalarNode creates a plain YAML scalar node.
func scalarNode(value string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Value: value}
}

// boolNode creates a YAML bool scalar node.
func boolNode(value bool) *yaml.Node {
	v := "false"
	if value {
		v = "true"
	}
	return &yaml.Node{Kind: yaml.ScalarNode, Value: v, Tag: "!!bool"}
}

func init() {
	// Register the same MCP flags as 'mcp serve' so users can configure them
	// with 'config set' without Viper binding conflicts.
	mcp.RegisterFlags(configSetCmd)
	Cmd.AddCommand(configSetCmd)
}
