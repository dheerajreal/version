package lib

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

type Tool struct {
	Name   string
	Binary string
	Args   []string
}

type ToolVersionResult struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

func (r ToolVersionResult) PrintToolVersionResult(){
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

var Tools = []Tool{
	// Self
	{"Version", "version", []string{"--version"}},

	// Languages
	{"Python", "python3", []string{"--version"}},
	{"Go", "go", []string{"version"}},

	// Package managers
	{"pip", "pip3", []string{"--version"}},
	{"uv", "uv", []string{"--version"}},

	// Dev tools
	{"Git", "git", []string{"--version"}},
	{"Docker", "docker", []string{"--version"}},
	{"kubectl", "kubectl", []string{"version", "--client"}},

}

// ─────────────────────────────────────────────────────────────
// Version detection
// ─────────────────────────────────────────────────────────────

var versionRe = regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?(?:[-+.]\w+)*)`)



func (t Tool) DetectToolVersion() ToolVersionResult {
	result := ToolVersionResult{Name: t.Name}
	path := which(t.Binary)
	if path == "" {
		return result
	}
	result.Path = path

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, t.Binary, t.Args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	_ = cmd.Run()

	lines := strings.Split(out.String(), "\n")
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			if m := versionRe.FindStringSubmatch(l); len(m) > 1 {
				result.Version = m[1]
				break
			} else {
				result.Version = l
			}
		}
	}

	return result
}
