package checker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	for _, t := range Tools {
		if strings.EqualFold(toolname, t.Name) || strings.EqualFold(toolname, t.Binary) {
			return t, nil
		}
	}
	return Tool{}, errors.New("Tool not supported")
}

var Tools = []Tool{
	// Self
	{"Version", "version", []string{"--version"}},

	// Languages
	{"Python", "python3", []string{"--version"}},
	{"Node.js", "node", []string{"--version"}},
	{"Go", "go", []string{"version"}},
	{"Rust", "rustc", []string{"--version"}},
	{"Java", "java", []string{"--version"}},
	{"Ruby", "ruby", []string{"--version"}},
	{"PHP", "php", []string{"-r", "echo phpversion();"}},
	{"Swift", "swift", []string{"--version"}},
	{"Kotlin", "kotlin", []string{"-version"}},
	{"Dart", "dart", []string{"--version"}},
	{"Perl", "perl", []string{"-v"}},
	{"Lua", "lua", []string{"-v"}},
	{"Elixir", "elixir", []string{"--version"}},
	{"R", "R", []string{"--version"}},
	{"Julia", "julia", []string{"--version"}},
	{"C / C++", "gcc", []string{"--version"}},
	{"Clang", "clang", []string{"--version"}},
	{"Zig", "zig", []string{"version"}},
	// Package managers
	{"npm", "npm", []string{"--version"}},
	{"yarn", "yarn", []string{"--version"}},
	{"pnpm", "pnpm", []string{"--version"}},
	{"bun", "bun", []string{"--version"}},
	{"pip", "pip3", []string{"--version"}},
	{"uv", "uv", []string{"--version"}},
	{"cargo", "cargo", []string{"--version"}},
	{"gem", "gem", []string{"--version"}},
	{"composer", "composer", []string{"--version"}},
	{"Homebrew", "brew", []string{"--version"}},
	{"luarocks", "luarocks", []string{"--version"}},
	{"opam", "opam", []string{"--version"}},
	{"pub", "pub", []string{"--version"}},
	// Dev tools
	{"Git", "git", []string{"--version"}},
	{"Docker", "docker", []string{"--version"}},
	{"kubectl", "kubectl", []string{"version", "--client"}},
	{"Helm", "helm", []string{"version"}},
	{"Terraform", "terraform", []string{"version"}},
	{"Ansible", "ansible", []string{"--version"}},
	{"Vagrant", "vagrant", []string{"--version"}},
	{"Make", "make", []string{"--version"}},
	{"CMake", "cmake", []string{"--version"}},
	{"Gradle", "gradle", []string{"--version"}},
	{"Maven", "mvn", []string{"--version"}},
	{"ESLint", "eslint", []string{"--version"}},
	{"Prettier", "prettier", []string{"--version"}},
	{"Black", "black", []string{"--version"}},
	{"Ruff", "ruff", []string{"--version"}},
	{"SQLite", "sqlite3", []string{"--version"}},
	{"jq", "jq", []string{"--version"}},
	{"curl", "curl", []string{"--version"}},
	// Cloud
	{"AWS CLI", "aws", []string{"--version"}},
	{"GCloud", "gcloud", []string{"version"}},
	{"Azure CLI", "az", []string{"version"}},
	{"Fly.io", "flyctl", []string{"version"}},
	{"Vercel", "vercel", []string{"--version"}},
	{"Netlify", "netlify", []string{"--version"}},
	{"Railway", "railway", []string{"--version"}},
	{"Render", "render", []string{"--version"}},
	{"Cloudflare", "wrangler", []string{"--version"}},
	{"Firebase", "firebase", []string{"--version"}},
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
	resultCh := make(chan ToolVersionResult, len(Tools))
	sem := make(chan struct{}, 5) // limit concurrency to 5 goroutines

	for _, t := range Tools {
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
