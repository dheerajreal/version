package checker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"sync"
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

func (r ToolVersionResult) PrintToolVersionResult() {
	path, version := r.Path, r.Version
	if path == "" || version == "" {
		fmt.Fprintf(os.Stderr, "Wrong fallback values")
		os.Exit(1)
	}
	fmt.Printf("%-15s %10s  %s\n", r.Name, version, path)
}

func FindTool(toolname string) (Tool, error) {
	for _, t := range toolsList {
		if strings.EqualFold(toolname, t.Name) || strings.EqualFold(toolname, t.Binary) {
			return t, nil
		}
	}
	return Tool{}, errors.New("Tool not supported")
}

// ─────────────────────────────────────────────────────────────
// Version detection
// ─────────────────────────────────────────────────────────────

var versionRe = regexp.MustCompile(`(\d+\.\d+(?:\.\d+)?(?:[-+.]\w+)*)`)

func (t Tool) DetectToolVersion() ToolVersionResult {
	result := ToolVersionResult{Name: t.Name}
	path := t.Where()
	if path == "" {
		result.Path = "Not Found"
		result.Version = "Unknown"
		return result
	}
	result.Path = path

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, t.Binary, t.Args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		result.Version = "Unknown"
		return result
	}

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

func (t Tool) Where() string {
	// return binary filepath
	bin := t.Binary
	path, err := exec.LookPath(bin)
	if err != nil {
		return ""
	}
	return path
}

func DetectAllToolsConcurrently() []ToolVersionResult {
	var wg sync.WaitGroup
	resultCh := make(chan ToolVersionResult, len(toolsList))
	// limit concurrency to (NumCPU/2) + 1 goroutines
	n := (runtime.NumCPU() / 2) + 1
	sem := make(chan struct{}, n)

	for _, t := range toolsList {
		wg.Add(1)
		go func(tool Tool) {
			defer wg.Done()
			sem <- struct{}{}        // acquire semaphore
			defer func() { <-sem }() // release semaphore

			result := tool.DetectToolVersion()
			resultCh <- result
		}(t)
	}

	wg.Wait()
	close(resultCh)

	var results []ToolVersionResult
	for r := range resultCh {
		results = append(results, r)
	}

	return results
}
