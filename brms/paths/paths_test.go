package paths

import (
	"path/filepath"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	// Utiliser filepath.Join pour construire le chemin de manière compatible avec tous les OS
	input := filepath.Join("..", "config", "examples", "subgroups.brms")
	result := GetFilePath(input)

	if result != input {
		t.Errorf("Expected %s, got %s", input, result)
	}
}
