package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	l "github.com/dheerajreal/version/lib"
)




func main() {
	args := os.Args[1:]
	jsonOutput := false
	showAll := false
	var toolName string

	for _, a := range args {
		switch a {
		case "--json":
			jsonOutput = true
		case "--all":
			showAll = true
		case "--version":
			// version --version
			fmt.Println(version)
			return
		default:
			toolName = a
		}
	}

	if !showAll && toolName == "" {
		fmt.Println(helpMessage)
		return
	}

	var results []l.ToolVersionResult

	if showAll {
		for _, t := range l.Tools {
			results = append(results, l.DetectTool(t))
		}
	} else if toolName != "" {
		found := false
		for _, t := range l.Tools {
			if strings.EqualFold(toolName, t.Name) || strings.EqualFold(toolName, t.Binary) {
				results = append(results, l.DetectTool(t))
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintf(os.Stderr, "Unknown tool: %s\n", toolName)
			os.Exit(1)
		}
	}

	if jsonOutput {
		data, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(data))

	} else {
		for _, r := range results {
			path := r.Path
			if path == "" {
				path = "not found"
			}
			version := r.Version
			if version == "" {
				version = "unknown"
			}
			fmt.Printf("%-15s %10s  %s\n", r.Name, version, path)
		}
	}
}



// ─────────────────────────────────────────────────────────────
// Help
// ─────────────────────────────────────────────────────────────

var helpMessage = `
Version: a version checker cli
Usage:

version                # prints this message
version --version      # Show version of version (meta)
version <toolname>     # Show a specific tool

Options:
	--json             # Output in JSON format
	--all              # Show all tools
`

// ─────────────────────────────────────────────────────────────
// version
// ─────────────────────────────────────────────────────────────

var (
	version = "0.0.1-dev"
)