//go:build windows
// +build windows

package paths

import "path/filepath"

// GetFilePath returns the correct path for Windows
func GetFilePath(path string) string {
	// Convert / to \\ and clean the path
	return filepath.Clean(filepath.FromSlash(path))
}
