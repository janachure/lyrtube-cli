package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

// Config holds settings for lyrtube-cli
type Config struct {
	PythonPath    string `mapstructure:"python_path"`
	YtdlpScript   string `mapstructure:"ytdlp_script"`
	OutputDir     string `mapstructure:"output_dir"`
	AudioFormat   string `mapstructure:"audio_format"`
	AudioQuality  string `mapstructure:"audio_quality"`
	LyricsEnabled bool   `mapstructure:"lyrics_enabled"`
}

var cfg *Config

// ConfigPath returns the path where the config lives
func ConfigPath() string {
	home, _ := os.UserHomeDir()
	confDir := filepath.Join(home, ".lyrtube-cli")
	_ = os.MkdirAll(confDir, 0o700)
	return filepath.Join(confDir, "config.yaml")
}

// EnsureConfig checks for existing config and runs wizard if missing
func EnsureConfig() error {
	path := ConfigPath()

	// STEP 1: check manually if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// STEP 2: run wizard
		c, err := RunSetupWizard()
		if err != nil {
			return fmt.Errorf("setup wizard failed: %w", err)
		}

		viper.SetConfigFile(path)
		viper.SetConfigType("yaml") // important

		viper.Set("python_path", c.PythonPath)
		viper.Set("yt_dlp_script", c.YtdlpScript)
		viper.Set("output_dir", c.OutputDir)
		viper.Set("audio_format", c.AudioFormat)
		viper.Set("audio_quality", c.AudioQuality)
		viper.Set("lyrics_enabled", c.LyricsEnabled)

		if err := viper.WriteConfigAs(path); err != nil {
			if os.IsNotExist(err) {
				if err := viper.SafeWriteConfigAs(path); err != nil {
					return fmt.Errorf("failed to write config: %w", err)
				}
			} else {
				return fmt.Errorf("failed to write config: %w", err)
			}
		}

		cfg = c
		return nil // IMPORTANT: stop here, don’t let viper.ReadInConfig() run
	}

	// STEP 3: load existing config
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	cfg = &c
	return nil
}

// Get returns the currently loaded config
func Get() *Config {
	if cfg == nil {
		_ = EnsureConfig()
	}
	return cfg
}

// RunSetupWizard asks the user for config interactively
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

	// Prompt helpers
	promptInput := func(label, def string) (string, error) {
		p := promptui.Prompt{
			Label:   label,
			Default: def,
		}
		return p.Run()
	}

	promptSelect := func(label string, items []string, def int) (string, error) {
		p := promptui.Select{
			Label:     label,
			Items:     items,
			Size:      len(items),
			CursorPos: def,
		}
		_, res, err := p.Run()
		return res, err
	}

	pythonPath, err := promptInput("Full path to Python interpreter (e.g., /usr/bin/python3)", "/usr/bin/python3")
	if err != nil {
		return nil, err
	}

	ytdlpPath, err := promptInput("Where is yt-dlp installed? (include also python3 if it's a python file)", "yt-dlp")
	if err != nil {
		return nil, err
	}

	home, _ := os.UserHomeDir()
	outputDir, err := promptInput("Where should we save downloads?", filepath.Join(home, "Download"))
	if err != nil {
		return nil, err
	}

	audioFormat, err := promptSelect("Default audio format", []string{"mp3", "m4a", "opus"}, 0)
	if err != nil {
		return nil, err
	}

	audioQuality, err := promptSelect("Default audio quality", []string{"128k", "192k", "320k"}, 0)
	if err != nil {
		return nil, err
	}

	lir, err := promptSelect("Enable automatic lyrics fetching?", []string{"Yes", "No"}, 0)
	if err != nil {
		return nil, err
	}

	pLyrics := (lir == "Yes")

	cfg := &Config{
		PythonPath:    pythonPath,
		YtdlpScript:   ytdlpPath,
		OutputDir:     outputDir,
		AudioFormat:   audioFormat,
		AudioQuality:  audioQuality,
		LyricsEnabled: pLyrics,
	}

	color.Green("\n✅ Configuration complete! It will be saved to %s\n", ConfigPath())
	return cfg, nil
}
