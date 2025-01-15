package paths

import (
	"runtime"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	input := "./../config/examples/subgroups.brms"

	expected := input
	if runtime.GOOS == "windows" {
		expected = "..\\config\\examples\\subgroups.brms"
	}

	result := GetFilePath(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
