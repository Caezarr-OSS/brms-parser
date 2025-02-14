//go:build !windows
// +build !windows

package paths

import "path/filepath"

// GetFilePath returns the correct path for Unix-like systems
func GetFilePath(path string) string {
	return filepath.Clean(path) // Cleans the path for Unix-like systems
}
