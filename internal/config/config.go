package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds settings for lyrtube-cli
type Config struct {
	OutputDir    string `mapstructure:"output_dir"`
	AudioFormat  string `mapstructure:"audio_format"`
	AudioQuality string `mapstructure:"audio_quality"`
	LyricsMode   string `mapstructure:"lyrics_mode"` // "none", "embedded", "lrc"
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

		viper.Set("output_dir", c.OutputDir)
		viper.Set("audio_format", c.AudioFormat)
		viper.Set("audio_quality", c.AudioQuality)
		viper.Set("lyrics_mode", c.LyricsMode)

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
