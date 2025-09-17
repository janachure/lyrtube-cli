package main

import (
	"fmt"
	"os"

	"lyrtube-cli/cmd"
	"lyrtube-cli/internal/config"

	"github.com/fatih/color"
)

func main() {
	// Ensure config exists before anything else
	if err := config.EnsureConfig(); err != nil {
		color.Red("Config error: %v", err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
