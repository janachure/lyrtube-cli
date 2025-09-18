package downloader

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"lyrtube-cli/internal/config"
	"lyrtube-cli/internal/embedded"
	"github.com/dhowden/tag"
)

// TrackMetadata holds metadata for a downloaded track
type TrackMetadata struct {
	Title    string
	Artist   string
	FilePath string
}

// DownloadAudio runs embedded yt-dlp with audio-only args and extracts metadata from file tags
func DownloadAudio(outDir, url, audioFormat, audioQuality string, cfg *config.Config) ([]TrackMetadata, error) {
	if strings.TrimSpace(url) == "" {
		return nil, fmt.Errorf("empty url")
	}

	if outDir == "" {
		outDir = cfg.OutputDir
	}
	if strings.HasPrefix(outDir, "~/") {
		home, _ := os.UserHomeDir()
		outDir = filepath.Join(home, outDir[2:])
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return nil, fmt.Errorf("failed to create output dir: %w", err)
	}

	// Download audio with metadata embedding
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

	cmd, err := embedded.RunYtDlp(args, outDir)
	if err != nil {
		return nil, fmt.Errorf("failed to create yt-dlp command: %w", err)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("yt-dlp failed: %w", err)
	}

	// Clean up temporary files
	defer embedded.CleanupTempFiles()

	// Find all downloaded audio files and extract metadata from each
	files, err := filepath.Glob(filepath.Join(outDir, "*."+audioFormat))
	if err != nil || len(files) == 0 {
		return nil, fmt.Errorf("no audio files found after download")
	}

	// Process all downloaded files
	var tracks []TrackMetadata
	for _, file := range files {
		// Extract metadata from audio file tags
		title, artist, err := extractMetadataFromFile(file)
		if err != nil {
			continue // Skip files with metadata extraction errors
		}

		tracks = append(tracks, TrackMetadata{
			Title:    title,
			Artist:   artist,
			FilePath: file,
		})
	}

	if len(tracks) == 0 {
		return nil, fmt.Errorf("no valid tracks found after download")
	}

	return tracks, nil
}

// extractMetadataFromFile reads audio file tags and returns cleaned title and artist
func extractMetadataFromFile(filePath string) (title, artist string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	m, err := tag.ReadFrom(file)
	if err != nil {
		return "", "", fmt.Errorf("failed to read tags: %w", err)
	}

	title = m.Title()
	artist = m.Artist()

	// Clean the title and artist for lyrics search
	title = cleanQueryString(title)
	artist = cleanQueryString(artist)

	return title, artist, nil
}

// cleanQueryString removes parentheses and their content, and cleans up the string
func cleanQueryString(s string) string {
	// Remove content in parentheses and brackets
	re := regexp.MustCompile(`\([^)]*\)|\[[^\]]*\]`)
	s = re.ReplaceAllString(s, "")
	
	// Remove extra whitespace
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	
	return s
}
