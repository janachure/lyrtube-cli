package downloader

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"lyrtube-cli/internal/config"
)

// DownloadAudio runs yt-dlp/yt-dlp wrapper with audio-only args
func DownloadAudio(pythonPath, ytdlpScript, outDir, url, audioFormat, audioQuality string, cfg *config.Config) (title, artist string, err error) {
	if strings.TrimSpace(url) == "" {
		return "", "", fmt.Errorf("empty url")
	}

	if outDir == "" {
		outDir = cfg.OutputDir
	}
	if strings.HasPrefix(outDir, "~/") {
		home, _ := os.UserHomeDir()
		outDir = filepath.Join(home, outDir[2:])
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return "", "", fmt.Errorf("failed to create output dir: %w", err)
	}

	// 1 Fetch metadata
	cmdMeta := exec.Command(pythonPath, ytdlpScript, "--get-title", "--get-artist", url)
	outMeta, err := cmdMeta.Output()
	if err != nil {
		return "", "", fmt.Errorf("failed to get metadata: %w", err)
	}
	lines := strings.Split(strings.TrimSpace(string(outMeta)), "\n")
	if len(lines) >= 2 {
		artist = lines[0]
		title = lines[1]
	} else if len(lines) == 1 {
		title = lines[0]
		artist = ""
	}

	// 2 Download audio
	args := []string{
		"--format", "bestaudio/best",
		"--extract-audio",
		"--audio-format", audioFormat,
		"--audio-quality", audioQuality,
		"--embed-metadata",
		"--embed-thumbnail",
		"--parse-metadata", "title:(?P<artist>.+?) - (?P<title>.+)",
		"--replace-in-metadata", "title", "(?P<artist>.+?) - ", "",
		"--output", "%(title)s.%(ext)s",
		"--ignore-errors",
		"--no-warnings",
		url,
	}

	cmd := exec.Command(pythonPath, append([]string{ytdlpScript}, args...)...)
	cmd.Dir = outDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", "", fmt.Errorf("yt-dlp failed: %w", err)
	}

	return title, artist, nil
}
