package cmd

import (
	"os"

	"lyrtube-cli/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lyrtube-cli",
	Short: "A small CLI to download audio using yt-dlp (audio-only)",
	Long:  "lyrtube-cli - download songs/audio using yt-dlp with a friendly terminal UI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Ensure config exists
		err := config.EnsureConfig()
		if err != nil {
			color.Red("Config error: %v", err)
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(configCmd)
}
