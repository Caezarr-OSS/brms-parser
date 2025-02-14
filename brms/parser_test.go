package brms

import (
	"testing"
	"os"
	"reflect"
)

func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "brms-test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	return tmpFile.Name()
}

func TestParser_Parse_ValidFile(t *testing.T) {
	content := `[block_a|block_b]
entity_1|entity_2
entity_3|entity_4

[block_c|block_d]
entity_5|entity_6`

	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	parser := NewParser(tmpFile, LogLevelInfo)
	parsed, err := parser.Parse()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check blocks
	expectedBlocks := map[string]string{
		"block_a": "block_b",
		"block_c": "block_d",
	}

	if !reflect.DeepEqual(parsed.Blocks, expectedBlocks) {
		t.Errorf("Expected blocks %v, got %v", expectedBlocks, parsed.Blocks)
	}

	// Check entities
	expectedEntities := []Mapping{
		{Source: "entity_1", Destination: "entity_2"},
		{Source: "entity_3", Destination: "entity_4"},
		{Source: "entity_5", Destination: "entity_6"},
	}

	if !reflect.DeepEqual(parsed.Entities, expectedEntities) {
		t.Errorf("Expected entities %v, got %v", expectedEntities, parsed.Entities)
	}
}

func TestParser_Parse_Exclusions(t *testing.T) {
	content := `[block_a|]
entity_1|
entity_2|

[block_b|]
entity_3|`

	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	parser := NewParser(tmpFile, LogLevelInfo)
	parsed, err := parser.Parse()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedExclusions := []Exclusion{
		{Source: "block_a"},
		{Source: "entity_1"},
		{Source: "entity_2"},
		{Source: "block_b"},
		{Source: "entity_3"},
	}

	if !reflect.DeepEqual(parsed.IgnoredItems, expectedExclusions) {
		t.Errorf("Expected exclusions %v, got %v", expectedExclusions, parsed.IgnoredItems)
	}
}

func TestParser_Parse_Subgroups(t *testing.T) {
	content := `[block_a/sub_1|block_b/sub_2]
entity_1|entity_2
entity_3|entity_4`

	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	parser := NewParser(tmpFile, LogLevelInfo)
	parsed, err := parser.Parse()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedBlocks := map[string]string{
		"block_a/sub_1": "block_b/sub_2",
	}

	if !reflect.DeepEqual(parsed.Blocks, expectedBlocks) {
		t.Errorf("Expected blocks %v, got %v", expectedBlocks, parsed.Blocks)
	}
}

func TestParser_Parse_CustomSeparator(t *testing.T) {
	content := `[block_a->block_b]
entity_1->entity_2
entity_3->entity_4`

	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile)

	parser := NewParser(tmpFile, LogLevelInfo)
	parser.SetSeparator("->")
	parsed, err := parser.Parse()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedBlocks := map[string]string{
		"block_a": "block_b",
	}

	if !reflect.DeepEqual(parsed.Blocks, expectedBlocks) {
		t.Errorf("Expected blocks %v, got %v", expectedBlocks, parsed.Blocks)
	}

	expectedEntities := []Mapping{
		{Source: "entity_1", Destination: "entity_2"},
		{Source: "entity_3", Destination: "entity_4"},
	}

	if !reflect.DeepEqual(parsed.Entities, expectedEntities) {
		t.Errorf("Expected entities %v, got %v", expectedEntities, parsed.Entities)
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
