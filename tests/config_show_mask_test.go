package tests

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"seerr-cli/cmd"

	"github.com/spf13/viper"
)

// writeConfigFile writes minimal YAML to a temp file so config show can read it.
func writeConfigFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, ".seerr-cli.yaml")
	if err := os.WriteFile(p, []byte(content), 0o600); err != nil {
		t.Fatalf("writeConfigFile: %v", err)
	}
	return p
}

func runConfigShow(t *testing.T, configPath string) string {
	t.Helper()
	viper.Reset()
	b := new(bytes.Buffer)
	cmd.RootCmd.SetOut(b)
	cmd.RootCmd.SetErr(b)
	cmd.RootCmd.SetArgs([]string{"config", "show", "--config", configPath})
	if err := cmd.RootCmd.Execute(); err != nil {
		t.Fatalf("config show: %v", err)
	}
	return b.String()
}

func TestConfigShowAPIKeyMasking(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		wantContain string
		wantAbsent  string
	}{
		{
			name:        "long key shows only last 4 chars",
			apiKey:      "abcdefghijklmnop",
			wantContain: "********mnop",
			wantAbsent:  "abcd",
		},
		{
			name:        "exactly 5 char key shows only last 4",
			apiKey:      "hello",
			wantContain: "****",
			// "hell" must not appear
			wantAbsent: "hell",
		},
		{
			name:        "4 char key is fully masked",
			apiKey:      "1234",
			wantContain: "****",
			wantAbsent:  "1234",
		},
		{
			name:        "short key (2 chars) is fully masked",
			apiKey:      "ab",
			wantContain: "****",
			wantAbsent:  "ab",
		},
		{
			name:        "key absent shows not-set label",
			apiKey:      "",
			wantContain: "<not set>",
			wantAbsent:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var yaml string
			if tc.apiKey != "" {
				yaml = "seerr:\n  server: http://localhost:5055\n  api_key: " + tc.apiKey + "\n"
			} else {
				yaml = "seerr:\n  server: http://localhost:5055\n"
			}
			p := writeConfigFile(t, yaml)
			out := runConfigShow(t, p)
			if !strings.Contains(out, tc.wantContain) {
				t.Errorf("expected output to contain %q, got: %s", tc.wantContain, out)
			}
			if tc.wantAbsent != "" && strings.Contains(out, tc.wantAbsent) {
				t.Errorf("expected output NOT to contain %q, got: %s", tc.wantAbsent, out)
			}
		})
	}
}
