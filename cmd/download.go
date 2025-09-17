package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"lyrtube-cli/internal/config"
	"lyrtube-cli/internal/downloader"
	"lyrtube-cli/internal/lyrics"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
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

		// CLI flags override config
		if outDir == "" {
			outDir = cfg.OutputDir
		}
		if fmtAudio == "" {
			fmtAudio = cfg.AudioFormat
		}
		if aq == "" {
			aq = cfg.AudioQuality
		}

		color.Yellow("Preparing to download audio from: %s", url)

		// start spinner
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Downloading..."
		s.Start()

		title, artist, err := downloader.DownloadAudio(cfg.PythonPath, cfg.YtdlpScript, outDir, url, fmtAudio, aq, cfg)
		s.Stop()

		if err != nil {
			color.Red("Download failed: %v", err)
			os.Exit(1)
		}

		color.Green("Download finished — saved to %s", outDir)

		if cfg.LyricsEnabled && title != "" {
			query := fmt.Sprintf("%s - %s", title, artist)
			color.Cyan("Searching lyrics for: %s", query)

			lr, lerr := lyrics.FetchLyrics(query)
			if lerr != nil {
				color.Yellow("Lyrics not found for \"%s\".", query)

				prompt := promptui.Prompt{
					Label:   "Enter alternative query for lyrics (or leave blank to skip)",
					Default: query,
				}
				alt, _ := prompt.Run()
				if strings.TrimSpace(alt) != "" {
					lr, lerr = lyrics.FetchLyrics(alt)
				}
			}

			if lerr == nil && lr != nil {
				lrcPath := filepath.Join(outDir, fmt.Sprintf("%s.lrc", title))
				if err := lyrics.WriteLRCFile(lrcPath, lr); err != nil {
					color.Red("Failed to save lyrics: %v", err)
				} else {
					color.Green("Lyrics saved to %s", lrcPath)
				}
			} else {
				color.Yellow("Skipping lyrics for \"%s\"", query)
			}
		}

	},
}

func init() {
	downloadCmd.Flags().StringVarP(&outDir, "out", "o", "", "Output directory (overrides config)")
	downloadCmd.Flags().StringVarP(&fmtAudio, "format", "f", "", "Audio format (mp3, m4a, etc) — overrides config")
	downloadCmd.Flags().StringVarP(&aq, "quality", "q", "", "Audio quality (e.g. 128k) — overrides config")
}
