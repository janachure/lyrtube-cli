package lyrics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type LRCLibResult struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	LRC    string `json:"syncedLyrics"`
}

// FetchLyrics queries LRCLIB for lyrics using title and artist
func FetchLyrics(title, artist string) (*LRCLibResult, error) {
	var url string
	if artist != "" {
		url = fmt.Sprintf("https://lrclib.net/api/get?artist_name=%s&track_name=%s",
			strings.ReplaceAll(artist, " ", "+"),
			strings.ReplaceAll(title, " ", "+"))
	} else {
		url = fmt.Sprintf("https://lrclib.net/api/get?track_name=%s",
			strings.ReplaceAll(title, " ", "+"))
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error contacting LRCLIB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("lyrics not found (status %d)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result LRCLibResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	if result.LRC == "" {
		return nil, fmt.Errorf("no synced lyrics found")
	}

	return &result, nil
}

// FetchLyricsByQuery queries LRCLIB for lyrics using a custom query string
func FetchLyricsByQuery(query string) (*LRCLibResult, error) {
	url := fmt.Sprintf("https://lrclib.net/api/get?track_name=%s", strings.ReplaceAll(query, " ", "+"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error contacting LRCLIB: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("lyrics not found (status %d)", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var result LRCLibResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	if result.LRC == "" {
		return nil, fmt.Errorf("no synced lyrics found")
	}

	return &result, nil
}

// WriteLRCFile saves the lyrics into a .lrc file
func WriteLRCFile(path string, lr *LRCLibResult) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	// Write metadata header
	if lr.Artist != "" {
		fmt.Fprintf(f, "[ar:%s]\n", lr.Artist)
	}
	if lr.Title != "" {
		fmt.Fprintf(f, "[ti:%s]\n", lr.Title)
	}

	// Lyrics are already time-synced from LRCLIB
	_, err = f.WriteString(lr.LRC)
	if err != nil {
		return fmt.Errorf("error writing lyrics: %w", err)
	}

	return nil
}

// EmbedLyricsIntoAudio embeds the LRC lyrics directly into the audio file metadata
func EmbedLyricsIntoAudio(audioPath string, lr *LRCLibResult) error {
	if lr == nil || lr.LRC == "" {
		return fmt.Errorf("no lyrics to embed")
	}

	// Get file extension and create proper temp file
	ext := filepath.Ext(audioPath)
	base := strings.TrimSuffix(audioPath, ext)
	tempPath := base + "_temp" + ext

	// Escape quotes and newlines in lyrics for shell command
	escapedLyrics := strings.ReplaceAll(lr.LRC, `"`, `\"`)
	escapedLyrics = strings.ReplaceAll(escapedLyrics, "\n", "\\n")
	escapedLyrics = strings.ReplaceAll(escapedLyrics, "$", "\\$")

	// Use ffmpeg to embed lyrics as metadata
	cmd := fmt.Sprintf(`ffmpeg -i "%s" -metadata lyrics="%s" -c copy "%s" -y`,
		audioPath,
		escapedLyrics,
		tempPath)

	// Execute ffmpeg command
	if err := executeCommand(cmd); err != nil {
		return fmt.Errorf("failed to embed lyrics: %w", err)
	}

	// Replace original file with the new one
	if err := os.Rename(tempPath, audioPath); err != nil {
		return fmt.Errorf("failed to replace original file: %w", err)
	}

	return nil
}

// executeCommand runs a shell command
func executeCommand(cmd string) error {
	execCmd := exec.Command("sh", "-c", cmd)
	output, err := execCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %s, output: %s", err, string(output))
	}
	return nil
}
