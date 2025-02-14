package main

import (
	"github.com/Caezarr-OSS/brms-parser/brms"
	"fmt"
	"log"
)

func main() {
	// Create a new parser with INFO log level
	parser := brms.NewParser("./config/example.brms", brms.LogLevelInfo)

	// Parse the BRMS file
	result, err := parser.Parse()
	if err != nil {
		log.Fatalf("Error while parsing: %v", err)
	}

	// Display mapped blocks
	fmt.Println("Mapped blocks:")
	for source, dest := range result.Blocks {
		fmt.Printf("  %s -> %s\n", source, dest)
	}

	// Display mapped entities
	fmt.Println("\nMapped entities:")
	for _, entity := range result.Entities {
		fmt.Printf("  %s -> %s\n", entity.Source, entity.Destination)
	}

	// Display ignored/excluded items
	fmt.Println("\nIgnored items:")
	for _, exclusion := range result.IgnoredItems {
		fmt.Printf("  %s\n", exclusion.Source)
	}

	// Example of handling exclusions
	processExclusions(result)
}

func processExclusions(result *brms.ParsedBRMS) {
	// Create a map for quick access to exclusions
	exclusions := make(map[string]bool)
	for _, excl := range result.IgnoredItems {
		exclusions[excl.Source] = true
	}

	// Example: Check if a specific project is excluded
	projectToCheck := "project_a"
	if exclusions[projectToCheck] {
		fmt.Printf("\nProject '%s' is excluded and won't be synchronized\n", projectToCheck)
	}

	// Example: Process all projects considering exclusions
	projects := []string{"project_a", "project_b", "project_c"}
	for _, project := range projects {
		if exclusions[project] {
			fmt.Printf("Exclusion: %s will not be processed\n", project)
		} else {
			fmt.Printf("Processing: %s will be synchronized\n", project)
		}
	}
}
