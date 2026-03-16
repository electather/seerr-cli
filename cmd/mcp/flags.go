package mcp

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FlagDef is the authoritative description of a single mcp-serve flag.
// Adding an entry here automatically propagates to both the serve command and
// the config set command — no other files need to be updated.
type FlagDef struct {
	// Name is the cobra flag name (e.g. "auth-token").
	Name string
	// ViperKey is the dot-notation Viper key (e.g. "mcp.auth_token").
	ViperKey string
	// Default is the zero/default value as a string (e.g. "stdio", "false").
	Default string
	// Usage is the help text shown by --help.
	Usage string
	// IsBool marks flags that are registered as Bool rather than String.
	IsBool bool
}

// ServeFlags is the single source of truth for all mcp serve flags.
// Order here determines the order of keys written to the config file.
var ServeFlags = []FlagDef{
	{
		Name:     "transport",
		ViperKey: "mcp.transport",
		Default:  "stdio",
		Usage:    "Transport protocol: stdio or http (env: SEERR_MCP_TRANSPORT)",
	},
	{
		Name:     "addr",
		ViperKey: "mcp.addr",
		Default:  ":8811",
		Usage:    "HTTP bind address (http transport only) (env: SEERR_MCP_ADDR)",
	},
	{
		Name:     "auth-token",
		ViperKey: "mcp.auth_token",
		Default:  "",
		Usage:    "Bearer token required for HTTP transport (env: SEERR_MCP_AUTH_TOKEN)",
	},
	{
		Name:     "no-auth",
		ViperKey: "mcp.no_auth",
		Default:  "false",
		IsBool:   true,
		Usage:    "Disable authentication (insecure — must be explicit) (env: SEERR_MCP_NO_AUTH)",
	},
	{
		Name:     "tls-cert",
		ViperKey: "mcp.tls_cert",
		Default:  "",
		Usage:    "Path to TLS certificate file (env: SEERR_MCP_TLS_CERT)",
	},
	{
		Name:     "tls-key",
		ViperKey: "mcp.tls_key",
		Default:  "",
		Usage:    "Path to TLS private key file (env: SEERR_MCP_TLS_KEY)",
	},
	{
		Name:     "cors",
		ViperKey: "mcp.cors",
		Default:  "false",
		IsBool:   true,
		Usage:    "Enable CORS headers (required for browser-based clients such as claude.ai) (env: SEERR_MCP_CORS)",
	},
	{
		Name:     "allow-api-key-query-param",
		ViperKey: "mcp.allow_api_key_query_param",
		Default:  "false",
		IsBool:   true,
		Usage:    "Accept the MCP auth token via the api_key query parameter in addition to headers (HTTP transport only)",
	},
	{
		Name:     "log-file",
		ViperKey: "mcp.log_file",
		Default:  "",
		Usage:    "Path to log file; required for stdio transport to capture logs (env: SEERR_MCP_LOG_FILE)",
	},
	{
		Name:     "log-level",
		ViperKey: "mcp.log_level",
		Default:  "info",
		Usage:    "Log level: debug, info, warn, error (env: SEERR_MCP_LOG_LEVEL)",
	},
	{
		Name:     "log-format",
		ViperKey: "mcp.log_format",
		Default:  "text",
		Usage:    "Log format: text or json (env: SEERR_MCP_LOG_FORMAT)",
	},
}

// RegisterFlags adds all MCP serve flags to cmd without binding them to Viper.
// Safe to call on any command (serve, config set, etc.) without side effects.
func RegisterFlags(cmd *cobra.Command) {
	for _, f := range ServeFlags {
		if f.IsBool {
			cmd.Flags().Bool(f.Name, f.Default == "true", f.Usage)
		} else {
			cmd.Flags().String(f.Name, f.Default, f.Usage)
		}
	}
}

// BindFlags binds the flags already registered on cmd to their Viper keys.
// Must only be called from the serve command so that config set does not
// overwrite the serve command's Viper bindings with unparsed flag values.
func BindFlags(cmd *cobra.Command) {
	for _, f := range ServeFlags {
		if flag := cmd.Flags().Lookup(f.Name); flag != nil {
			viper.BindPFlag(f.ViperKey, flag)
		}
	}
}
