package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

// Spinner wraps the spinner functionality with convenient methods
type Spinner struct {
	s *spinner.Spinner
}

// NewSpinner creates a new spinner with default settings
func NewSpinner(message string) *Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = " " + message
	return &Spinner{s: s}
}

// Start begins the spinner animation
func (sp *Spinner) Start() {
	sp.s.Start()
}

// Stop ends the spinner animation
func (sp *Spinner) Stop() {
	sp.s.Stop()
}
