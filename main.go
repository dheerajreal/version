package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	checker "github.com/dheerajreal/version/lib"
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

	if showAll && toolName != "" {
		fmt.Println("cannot pass --all with tool name")
		os.Exit(1)
	}

	var results []checker.ToolVersionResult

	if showAll {
		for _, t := range checker.Tools {
			results = append(results, t.DetectToolVersion())
		}
	} else if toolName != "" {
		found := false
		for _, t := range checker.Tools {
			if strings.EqualFold(toolName, t.Name) || strings.EqualFold(toolName, t.Binary) {
				results = append(results, t.DetectToolVersion())
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
			r.PrintToolVersionResult()
		}
	}
}



// ─────────────────────────────────────────────────────────────
// Help
// ─────────────────────────────────────────────────────────────

// made with `figlet -f future version``
var helpMessage = `
╻ ╻┏━╸┏━┓┏━┓╻┏━┓┏┓╻
┃┏┛┣╸ ┣┳┛┗━┓┃┃ ┃┃┗┫
┗┛ ┗━╸╹┗╸┗━┛╹┗━┛╹ ╹

A version checker cli
Usage:

version                # prints this message
version --version      # Show version of version (meta)
version <toolname>     # Show version of a specific tool

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