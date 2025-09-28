package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"stereo-server/internal/cli/process"
)

var rootCmd = &cobra.Command{
	Use:   "stereo-server",
	Short: "CLI for process JSON files for Stereo",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cli_process.NewCommand())
}
