package main

import (
	"seerr-cli/cmd"
	"seerr-cli/cmd/mcp"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	mcp.SetVersionInfo(version)
	cmd.Execute()
}
