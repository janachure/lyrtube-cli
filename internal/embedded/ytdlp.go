//go:build !windows
// +build !windows

package embedded

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

//go:embed yt-dlp
var ytdlpBinary []byte

// GetYtDlpPath extracts the embedded yt-dlp binary to a temporary location and returns the path
func GetYtDlpPath() (string, error) {
	// Create a temporary directory for the binary
	tempDir, err := os.MkdirTemp("", "lyrtube-ytdlp-*")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	// Determine the binary name based on OS
	binaryName := "yt-dlp"
	if runtime.GOOS == "windows" {
		binaryName = "yt-dlp.exe"
	}

	binaryPath := filepath.Join(tempDir, binaryName)

	// Write the embedded binary to the temp file
	err = os.WriteFile(binaryPath, ytdlpBinary, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to write yt-dlp binary: %w", err)
	}

	return binaryPath, nil
}

// RunYtDlp executes yt-dlp with the given arguments
func RunYtDlp(args []string, workDir string) (*exec.Cmd, error) {
	binaryPath, err := GetYtDlpPath()
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("python3", append([]string{binaryPath}, args...)...)
	if workDir != "" {
		cmd.Dir = workDir
	}

	return cmd, nil
}

// RunYtDlpOutput executes yt-dlp and returns the output
func RunYtDlpOutput(args []string, workDir string) ([]byte, error) {
	cmd, err := RunYtDlp(args, workDir)
	if err != nil {
		return nil, err
	}

	return cmd.Output()
}

// CleanupTempFiles removes temporary yt-dlp files (call this when done)
func CleanupTempFiles() {
	// Find and remove temp directories created by this package
	tempDir := os.TempDir()
	entries, err := os.ReadDir(tempDir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() && len(entry.Name()) > 15 && entry.Name()[:15] == "lyrtube-ytdlp-" {
			fullPath := filepath.Join(tempDir, entry.Name())
			os.RemoveAll(fullPath)
		}
	}
}