package main

import (
	"encoding/json"
	"fmt"
	"os"

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

	if !showAll && toolName == "" || toolName == "--help" {
		fmt.Println(helpMessage)
		return
	}

	if showAll && toolName != "" {
		fmt.Println("cannot pass --all with tool name")
		os.Exit(1)
	}

	var results []checker.ToolVersionResult

	if showAll {
		results = checker.DetectAllToolsConcurrently()
	} else if toolName != "" {
		tool, err := checker.FindTool(toolName)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+": %s\n", toolName)
			os.Exit(1)
		}
		results = append(results, tool.DetectToolVersion())
	}

	if jsonOutput {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to convert to json")
			os.Exit(1)
		}
		fmt.Println(string(data))

	} else {
		for _, r := range results {
			if r.Path != "Not Found" || !showAll {
				// print only the ones that are found in PATH
				r.PrintToolVersionResult()
			}
		}
	}
}

// ─────────────────────────────────────────────────────────────
// Help
// ─────────────────────────────────────────────────────────────

// made with `figlet -f future version“
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
