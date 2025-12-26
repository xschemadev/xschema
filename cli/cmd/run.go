package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Parse codebase, convert schemas, output native validators",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run: not implemented")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
