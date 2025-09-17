package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"lyrtube-cli/internal/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage lyrtube-cli configuration",
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the config file in $EDITOR (or nano)",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.ConfigPath()
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		c := exec.Command(editor, path)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			color.Red("Failed to open editor: %v", err)
			os.Exit(1)
		}
		color.Green("Saved config: %s", path)
	},
}

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the active config",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()
		fmt.Printf("Config file: %s\n", config.ConfigPath())
		fmt.Printf("pthon path: %s\n", cfg.PythonPath)
		fmt.Printf("yt-dlp path: %s\n", cfg.YtdlpScript)
		fmt.Printf("output dir: %s\n", cfg.OutputDir)
		fmt.Printf("audio fmt: %s\n", cfg.AudioFormat)
		fmt.Printf("audio quality: %s\n", cfg.AudioQuality)
	},
}

func init() {
	configCmd.AddCommand(configEditCmd)
	configCmd.AddCommand(configShowCmd)
}
