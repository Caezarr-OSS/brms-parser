package main

import (
	"github.com/Caezarr-OSS/brms-parser/brms"
	"github.com/Caezarr-OSS/brms-parser/brms/paths"
	"fmt"
	"log"
)

func main() {
	// Relative path to the BRMS configuration file
	relativePath := "./config/examples/subgroups.brms"

	// Dynamically convert to an OS-specific absolute path
	filePath := paths.GetFilePath(relativePath)

	// Print the resolved file path
	fmt.Printf("Using file path: %s\n", filePath)

	// Initialize the parser with the specified file and log level
	parser := brms.NewParser(filePath, brms.LogLevelInfo)

	// Parse the BRMS file and retrieve structured data
	parsed, err := parser.Parse()
	if err != nil {
		log.Fatalf("Error while parsing file: %v\n", err)
	}

	// Display parsed blocks
	fmt.Println("Blocks:")
	for source, dest := range parsed.Blocks {
		fmt.Printf("  %s -> %s\n", source, dest)
	}

	// Display parsed entities
	fmt.Println("\nEntities:")
	for _, entity := range parsed.Entities {
		fmt.Printf("  %s -> %s\n", entity.Source, entity.Destination)
	}

	// Display parsed exclusions
	fmt.Println("\nExclusions:")
	for _, exclusion := range parsed.IgnoredItems {
		fmt.Printf("  %s\n", exclusion.Source)
	}
}
