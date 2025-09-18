package cmd

import (
	"errors"
	"os"

	"lyrtube-cli/internal/config"
	"lyrtube-cli/internal/downloader"
	"lyrtube-cli/internal/processor"
	"lyrtube-cli/internal/ui"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	// flags
	outDir   string
	fmtAudio string
	aq       string
)

var downloadCmd = &cobra.Command{
	Use:   "download <url>",
	Short: "Download audio from a URL (audio only)",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires exactly one URL argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		cfg := config.Get()

		// Apply CLI flag overrides
		outputDir := getOutputDir(cfg)
		audioFormat := getAudioFormat(cfg)
		audioQuality := getAudioQuality(cfg)

		color.Yellow("Preparing to download audio from: %s", url)

		// Download with spinner
		spinner := ui.NewSpinner("Downloading...")
		spinner.Start()

		tracks, err := downloader.DownloadAudio(outputDir, url, audioFormat, audioQuality, cfg)
		if err != nil {
			spinner.Stop()
			color.Red("Download failed: %v", err)
			os.Exit(1)
		}
		spinner.Stop()
		color.Green("Download finished — %d track(s) saved to %s", len(tracks), outputDir)

		// Process lyrics for each track if needed
		lyricsProcessor := processor.NewLyricsProcessor(cfg)
		for i, track := range tracks {
			if len(tracks) > 1 {
				color.Cyan("Processing lyrics for track %d/%d: %s", i+1, len(tracks), track.Title)
			}
			if err := lyricsProcessor.ProcessLyrics(track.FilePath, track.Title, track.Artist, audioFormat); err != nil {
				color.Yellow("Lyrics processing completed with warnings for: %s", track.Title)
			}
		}
	},
}

// Helper functions for CLI flag overrides
func getOutputDir(cfg *config.Config) string {
	if outDir != "" {
		return outDir
	}
	return cfg.OutputDir
}

func getAudioFormat(cfg *config.Config) string {
	if fmtAudio != "" {
		return fmtAudio
	}
	return cfg.AudioFormat
}

func getAudioQuality(cfg *config.Config) string {
	if aq != "" {
		return aq
	}
	return cfg.AudioQuality
}

func init() {
	downloadCmd.Flags().StringVarP(&outDir, "out", "o", "", "Output directory (overrides config)")
	downloadCmd.Flags().StringVarP(&fmtAudio, "format", "f", "", "Audio format (mp3, m4a, etc) — overrides config")
	downloadCmd.Flags().StringVarP(&aq, "quality", "q", "", "Audio quality (e.g. 128k) — overrides config")
}
