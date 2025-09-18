package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// RunSetupWizard guides the user through initial configuration
func RunSetupWizard() (*Config, error) {
	banner := color.CyanString(`
██╗  ██╗   ██╗██████╗ ████████╗██╗   ██╗██████╗ ███████╗
██║  ╚██╗ ██╔╝██╔══██╗╚══██╔══╝██║   ██║██╔══██╗██╔════╝
██║   ╚████╔╝ ██████╔╝   ██║   ██║   ██║██████╔╝█████╗  
██║    ╚██╔╝  ██╔══██╗   ██║   ██║   ██║██╔══██╗██╔══╝  
███████╗██║   ██║  ██║   ██║   ╚██████╔╝██████╔╝███████╗
╚══════╝╚═╝   ╚═╝  ╚═╝   ╚═╝    ╚═════╝ ╚═════╝ ╚══════╝
   🎵 Welcome to lyrtube-cli! 🎵`)
	fmt.Println(banner)
	fmt.Println("It looks like this is your first time running the app. Let's configure it together!")

	// Output directory
	outputDir, err := promptOutputDirectory()
	if err != nil {
		return nil, err
	}

	// Audio format
	audioFormat, err := promptAudioFormat()
	if err != nil {
		return nil, err
	}

	// Audio quality
	audioQuality, err := promptAudioQuality()
	if err != nil {
		return nil, err
	}

	// Lyrics mode
	lyricsMode, err := promptLyricsMode()
	if err != nil {
		return nil, err
	}

	config := &Config{
		OutputDir:    outputDir,
		AudioFormat:  audioFormat,
		AudioQuality: audioQuality,
		LyricsMode:   lyricsMode,
	}

	color.Green("Configuration saved successfully!")
	return config, nil
}

// promptOutputDirectory asks user for download directory
func promptOutputDirectory() (string, error) {
	home, _ := os.UserHomeDir()
	defaultDir := filepath.Join(home, "Downloads")

	prompt := promptui.Prompt{
		Label:   "Download location",
		Default: defaultDir,
		Validate: func(input string) error {
			if strings.TrimSpace(input) == "" {
				return fmt.Errorf("download location cannot be empty")
			}
			return nil
		},
	}

	return prompt.Run()
}

// promptAudioFormat asks user for preferred audio format
func promptAudioFormat() (string, error) {
	prompt := promptui.Select{
		Label: "Audio format",
		Items: []string{"mp3", "m4a", "flac", "wav"},
	}

	_, result, err := prompt.Run()
	return result, err
}

// promptAudioQuality asks user for preferred audio quality
func promptAudioQuality() (string, error) {
	prompt := promptui.Select{
		Label: "Audio quality",
		Items: []string{"128k", "192k", "256k", "320k"},
	}

	_, result, err := prompt.Run()
	return result, err
}

// promptLyricsMode asks user for lyrics handling preference
func promptLyricsMode() (string, error) {
	prompt := promptui.Select{
		Label: "Lyrics mode",
		Items: []string{
			"embedded (embed lyrics into audio file)",
			"lrc (separate .lrc file)",
			"none (no lyrics)",
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	// Extract the mode from the descriptive text
	switch {
	case strings.HasPrefix(result, "embedded"):
		return "embedded", nil
	case strings.HasPrefix(result, "lrc"):
		return "lrc", nil
	case strings.HasPrefix(result, "none"):
		return "none", nil
	default:
		return "embedded", nil // fallback
	}
}