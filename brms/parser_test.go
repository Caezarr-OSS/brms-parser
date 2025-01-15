package brms

import (
	"testing"
)

func TestParser_Parse_ValidFile(t *testing.T) {
	tempFile := "./../config/examples/valid.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.Blocks) != 2 {
		t.Errorf("Expected 2 blocks, got %d", len(parsed.Blocks))
	}

	if len(parsed.Entities) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(parsed.Entities))
	}
}

func TestParser_Parse_Exclusions(t *testing.T) {
	tempFile := "./../config/examples/exclusions.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.IgnoredItems) != 2 {
		t.Errorf("Expected 2 exclusions, got %d", len(parsed.IgnoredItems))
	}

	if len(parsed.Entities) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(parsed.Entities))
	}
}

func TestParser_Parse_InvalidFile(t *testing.T) {
	tempFile := "./../config/examples/invalid.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	_, err := parser.Parse()
	if err == nil {
		t.Error("Expected parsing to fail for invalid file, but it succeeded")
	}
}

func TestParser_Parse_CommentsAndEmptyLines(t *testing.T) {
	tempFile := "./../config/examples/comments_and_empty.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.Blocks) != 1 {
		t.Errorf("Expected 1 block, got %d", len(parsed.Blocks))
	}

	if len(parsed.Entities) != 1 {
		t.Errorf("Expected 1 entity, got %d", len(parsed.Entities))
	}
}

func TestParser_Parse_EmptyFile(t *testing.T) {
	tempFile := "./../config/examples/empty.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.Blocks) != 0 || len(parsed.Entities) != 0 || len(parsed.IgnoredItems) != 0 {
		t.Error("Expected no blocks, entities, or exclusions in empty file")
	}
}

func TestParser_Parse_IndentationWarnings(t *testing.T) {
	tempFile := "./../config/examples/indentation.brms"
	parser := NewParser(tempFile, LogLevelWarn)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.Blocks) != 1 {
		t.Errorf("Expected 1 block, got %d", len(parsed.Blocks))
	}

	if len(parsed.Entities) != 2 {
		t.Errorf("Expected 2 entities, got %d", len(parsed.Entities))
	}
}

func TestParser_Parse_Subgroups(t *testing.T) {
	tempFile := "./../config/examples/subgroups.brms"
	parser := NewParser(tempFile, LogLevelInfo)

	parsed, err := parser.Parse()
	if err != nil {
		t.Fatalf("Parsing failed: %v", err)
	}

	if len(parsed.Blocks) != 2 {
		t.Errorf("Expected 2 blocks, got %d", len(parsed.Blocks))
	}

	if len(parsed.Entities) != 3 {
		t.Errorf("Expected 3 entities, got %d", len(parsed.Entities))
	}
}
