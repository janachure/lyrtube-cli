package lyrics

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type LRCLibResult struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
	LRC    string `json:"syncedLyrics"`
}

// FetchLyrics queries LRCLIB for lyrics given a search string
func FetchLyrics(query string) (*LRCLibResult, error) {
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
