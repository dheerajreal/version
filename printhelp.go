package main

import "fmt"

// ─────────────────────────────────────────────────────────────
// Help
// ─────────────────────────────────────────────────────────────

func printHelp() {
	fmt.Println(
	`
	Version: a version checker cli
	Usage:
		version                # prints this message
		version --version      # Show version of version (meta)
		version <toolname>     # Show a specific tool

		Options:
			--json             # Output in JSON format
			--all              # Show all tools
		`,
	)
}
