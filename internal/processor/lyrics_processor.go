package processor

import (
	"fmt"
	"path/filepath"
	"strings"

	"lyrtube-cli/internal/config"
	"lyrtube-cli/internal/lyrics"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

// LyricsProcessor handles lyrics fetching and processing
type LyricsProcessor struct {
	cfg *config.Config
}

// NewLyricsProcessor creates a new lyrics processor
func NewLyricsProcessor(cfg *config.Config) *LyricsProcessor {
	return &LyricsProcessor{cfg: cfg}
}

// ProcessLyrics handles the complete lyrics workflow for a downloaded track
func (lp *LyricsProcessor) ProcessLyrics(audioFilePath, title, artist, audioFormat string) error {
	if lp.cfg.LyricsMode == "none" || title == "" {
		return nil
	}

	// Search for lyrics
	lr, err := lp.searchLyrics(title, artist)
	if err != nil {
		return nil // Not a fatal error, just skip lyrics
	}

	// Create safe filename
	safeTitle := lp.createSafeFilename(title)

	// Handle lyrics based on config preference
	return lp.handleLyricsMode(audioFilePath, safeTitle, audioFormat, lr)
}

// searchLyrics attempts to find lyrics with fallback to user input
func (lp *LyricsProcessor) searchLyrics(title, artist string) (*lyrics.LRCLibResult, error) {
	var searchTitle, searchArtist string
	if artist != "" {
		color.Cyan("Searching lyrics for: %s - %s", artist, title)
		searchTitle, searchArtist = title, artist
	} else {
		color.Cyan("Searching lyrics for: %s", title)
		searchTitle, searchArtist = title, ""
	}

	lr, err := lyrics.FetchLyrics(searchTitle, searchArtist)
	if err != nil {
		return lp.handleLyricsNotFound(searchTitle, searchArtist)
	}

	return lr, nil
}

// handleLyricsNotFound prompts user for alternative search or skip
func (lp *LyricsProcessor) handleLyricsNotFound(searchTitle, searchArtist string) (*lyrics.LRCLibResult, error) {
	query := fmt.Sprintf("%s - %s", searchArtist, searchTitle)
	if searchArtist == "" {
		query = searchTitle
	}
	color.Yellow("Lyrics not found for \"%s\".", query)

	// Ask user what to do
	prompt := promptui.Select{
		Label: "What would you like to do?",
		Items: []string{"Skip lyrics", "Try different search query"},
	}
	_, choice, err := prompt.Run()
	if err != nil || choice == "Skip lyrics" {
		color.Yellow("Skipping lyrics for \"%s\"", query)
		return nil, fmt.Errorf("user skipped lyrics")
	}

	// Let user enter custom query
	queryPrompt := promptui.Prompt{
		Label:   "Enter search query for lyrics",
		Default: query,
	}
	alt, err := queryPrompt.Run()
	if err != nil || strings.TrimSpace(alt) == "" {
		color.Yellow("Skipping lyrics.")
		return nil, fmt.Errorf("user skipped lyrics")
	}

	lr, err := lyrics.FetchLyricsByQuery(strings.TrimSpace(alt))
	if err != nil {
		color.Yellow("Still no lyrics found. Skipping.")
		return nil, err
	}

	return lr, nil
}

// createSafeFilename creates a filesystem-safe filename from title
func (lp *LyricsProcessor) createSafeFilename(title string) string {
	safeTitle := strings.ReplaceAll(title, "/", "-")
	safeTitle = strings.ReplaceAll(safeTitle, ":", "-")
	safeTitle = strings.ReplaceAll(safeTitle, "?", "")
	safeTitle = strings.ReplaceAll(safeTitle, "*", "")
	safeTitle = strings.ReplaceAll(safeTitle, "<", "")
	safeTitle = strings.ReplaceAll(safeTitle, ">", "")
	safeTitle = strings.ReplaceAll(safeTitle, "|", "-")
	safeTitle = strings.ReplaceAll(safeTitle, "\\", "-")
	return safeTitle
}

// handleLyricsMode processes lyrics according to the configured mode
func (lp *LyricsProcessor) handleLyricsMode(audioFilePath, safeTitle, audioFormat string, lr *lyrics.LRCLibResult) error {
	switch lp.cfg.LyricsMode {
	case "embedded":
		// Embed lyrics directly into audio file
		audioPath := audioFilePath
		if err := lyrics.EmbedLyricsIntoAudio(audioPath, lr); err != nil {
			color.Yellow("Warning: Failed to embed lyrics into audio file: %v", err)
			// Fallback to .lrc file
			dir := filepath.Dir(audioFilePath)
			lrcPath := filepath.Join(dir, fmt.Sprintf("%s.lrc", safeTitle))
			if err := lyrics.WriteLRCFile(lrcPath, lr); err != nil {
				color.Red("Failed to save lyrics: %v", err)
				return err
			}
			color.Green("Lyrics saved to %s (fallback)", lrcPath)
		} else {
			color.Green("Lyrics embedded into audio file")
		}
	case "lrc":
		// Save lyrics to .lrc file
		dir := filepath.Dir(audioFilePath)
		lrcPath := filepath.Join(dir, fmt.Sprintf("%s.lrc", safeTitle))
		if err := lyrics.WriteLRCFile(lrcPath, lr); err != nil {
			color.Red("Failed to save lyrics: %v", err)
			return err
		}
		color.Green("Lyrics saved to %s", lrcPath)
	}
	return nil
}
