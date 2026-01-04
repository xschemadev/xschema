package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

// Version information set by goreleaser
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "xschema",
	Short: "JSON Schema to native validators",
}

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("xschema {{.Version}} (" + commit + ", " + date + ")\n")
}
