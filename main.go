package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/dheerajreal/version/lib/checker"
)

func main() {
	jsonOutput := flag.Bool("json", false, "Show output in JSON format")
	showAll := flag.Bool("all", false, "Show all tools")
	showVersion := flag.Bool("version", false, "Show version of version")

	flag.Usage = func() {
		fmt.Print(helpMessage)
	}

	flag.Parse()

	// version --version
	if *showVersion {
		fmt.Println(version)
		return
	}

	// positional arguments
	args := flag.Args()

	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		os.Exit(1)
	}

	var toolName string

	if len(args) == 1 {
		toolName = args[0]
	}

	handleCommand(*jsonOutput, *showAll, toolName)
}

func handleCommand(jsonOutput bool, showAll bool, toolName string){
	if !showAll && toolName == "" {
		flag.Usage()
		return
	}

	if showAll && toolName != "" {
		fmt.Fprintln(os.Stderr, "cannot pass --all with tool name")
		os.Exit(1)
	}

	var results []checker.ToolVersionResult

	if showAll {
		results = checker.DetectAllToolsConcurrently()
	} else if toolName != "" {
		tool, err := checker.FindTool(toolName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v: %s\n", err, toolName)
			os.Exit(1)
		}
		results = append(results, tool.DetectToolVersion())
	}

	if jsonOutput {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintln(os.Stderr, "failed to convert to json")
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

// made with `figlet -f future version`
var helpMessage = `
╻ ╻┏━╸┏━┓┏━┓╻┏━┓┏┓╻
┃┏┛┣╸ ┣┳┛┗━┓┃┃ ┃┃┗┫
┗┛ ┗━╸╹┗╸┗━┛╹┗━┛╹ ╹

A version checker CLI
Usage:

version [options] <toolname>

Examples:

version                # prints help
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
	version = "v0.0.2-dev"
)
