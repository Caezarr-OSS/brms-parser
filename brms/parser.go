package brms

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LogLevel defines the verbosity of logs.
const (
	LogLevelInfo  = "INFO"
	LogLevelWarn  = "WARN"
	LogLevelError = "ERROR"
)

// BRMSError represents an error specific to parsing BRMS files.
type BRMSError struct {
	Line    int    // Line number where the error occurred.
	Message string // Detailed error message.
}

func (e *BRMSError) Error() string {
	return fmt.Sprintf("Error at line %d: %s", e.Line, e.Message)
}

// Mapping represents a source-destination relationship.
type Mapping struct {
	Source      string // The source block or entity.
	Destination string // The destination block or entity.
}

// Exclusion represents a block or entity explicitly excluded.
type Exclusion struct {
	Source string // The excluded block or entity.
}

// ParsedBRMS holds the parsed data from a BRMS configuration file.
type ParsedBRMS struct {
	Blocks       map[string]string // Maps source blocks to destination blocks.
	Entities     []Mapping         // Lists entity mappings from source to destination.
	IgnoredItems []Exclusion       // Blocks or entities explicitly excluded.
}

// Parser handles parsing of a BRMS configuration file.
type Parser struct {
	FilePath  string // Path to the BRMS file to parse.
	LogLevel  string // Logging verbosity level.
	Separator string // Custom separator for mappings (default "|")
}

// NewParser creates a new Parser for the given file path.
func NewParser(filePath string, logLevel string) *Parser {
	return &Parser{
		FilePath:  filePath,
		LogLevel:  logLevel,
		Separator: "|", // Default separator changed to |
	}
}

// SetSeparator allows changing the mapping separator
func (p *Parser) SetSeparator(sep string) {
	p.Separator = sep
}

// log outputs messages based on the log level.
func (p *Parser) log(level string, message string) {
	if level == LogLevelError ||
		(level == LogLevelWarn && (p.LogLevel == LogLevelInfo || p.LogLevel == LogLevelWarn)) ||
		(level == LogLevelInfo && p.LogLevel == LogLevelInfo) {
		fmt.Printf("[%s] %s\n", level, message)
	}
}

// Parse reads and interprets a BRMS configuration file.
func (p *Parser) Parse() (*ParsedBRMS, error) {
	file, err := os.Open(p.FilePath)
	if err != nil {
		return nil, &BRMSError{Line: 0, Message: "failed to open file: " + err.Error()}
	}
	defer file.Close()

	parsed := &ParsedBRMS{
		Blocks:       make(map[string]string),
		Entities:     []Mapping{},
		IgnoredItems: []Exclusion{},
	}

	scanner := bufio.NewScanner(file)
	var currentBlock string
	var isExclusionBlock bool // Tracks if the current block is an exclusion
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		// Check for excessive spaces or tabs.
		if strings.TrimSpace(line) != line {
			p.log(LogLevelWarn, fmt.Sprintf("Line %d: Excessive spaces or tabs detected", lineNumber))
		}

		line = strings.TrimSpace(line)

		// Ignore empty lines or comments.
		if line == "" || strings.HasPrefix(line, "#") {
			p.log(LogLevelInfo, fmt.Sprintf("Line %d: Ignoring empty/comment line", lineNumber))
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// Handle a new block section.
			blockLine := strings.Trim(line, "[]")
			parts := strings.Split(blockLine, p.Separator)
			if len(parts) > 2 {
				return nil, &BRMSError{Line: lineNumber, Message: fmt.Sprintf("invalid format with multiple '%s' in block section", p.Separator)}
			}
			if len(parts) == 1 || parts[1] == "" {
				// Block exclusion case.
				parsed.IgnoredItems = append(parsed.IgnoredItems, Exclusion{Source: parts[0]})
				isExclusionBlock = true
				p.log(LogLevelInfo, fmt.Sprintf("Line %d: Parsed block exclusion '%s'", lineNumber, parts[0]))
				currentBlock = ""
			} else {
				// Valid block mapping.
				parsed.Blocks[parts[0]] = parts[1]
				isExclusionBlock = false
				currentBlock = blockLine
				p.log(LogLevelInfo, fmt.Sprintf("Line %d: Parsed block '%s'", lineNumber, blockLine))
			}
		} else if currentBlock != "" || isExclusionBlock {
			// Handle an entity mapping or exclusion.
			// Remove inline comments first
			if idx := strings.Index(line, "#"); idx != -1 {
				line = strings.TrimSpace(line[:idx])
			}

			parts := strings.Split(line, p.Separator)
			if len(parts) > 2 {
				return nil, &BRMSError{Line: lineNumber, Message: fmt.Sprintf("invalid format with multiple '%s' in entity mapping", p.Separator)}
			}
			if len(parts) == 1 || parts[1] == "" {
				// Exclusion case.
				parsed.IgnoredItems = append(parsed.IgnoredItems, Exclusion{Source: parts[0]})
				p.log(LogLevelInfo, fmt.Sprintf("Line %d: Parsed exclusion for '%s'", lineNumber, parts[0]))
			} else {
				// Entity mapping case.
				parsed.Entities = append(parsed.Entities, Mapping{
					Source:      parts[0],
					Destination: parts[1],
				})
				p.log(LogLevelInfo, fmt.Sprintf("Line %d: Parsed entity '%s' under block '%s'", lineNumber, line, currentBlock))
			}
		} else {
			// Line outside of a block section.
			p.log(LogLevelError, fmt.Sprintf("Line %d: Line '%s' is outside of any block section", lineNumber, line))
			return nil, &BRMSError{Line: lineNumber, Message: "line outside of block section"}
		}

		// Check for suspicious indentation (e.g., leading spaces or tabs).
		if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
			p.log(LogLevelWarn, fmt.Sprintf("Line %d: Suspicious indentation detected", lineNumber))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, &BRMSError{Line: lineNumber, Message: "failed to read file: " + err.Error()}
	}

	return parsed, nil // Successfully parsed the file.
}
