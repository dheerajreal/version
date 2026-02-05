package checker

var toolsList = []Tool{
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

	// Python ecosystem
	{"pipx", "pipx", []string{"--version"}},
	{"poetry", "poetry", []string{"--version"}},

	// Node / JS ecosystem
	{"Deno", "deno", []string{"--version"}},
	{"TypeScript", "tsc", []string{"--version"}},
	{"nvm", "nvm", []string{"--version"}},

	// Git-related
	{"Git LFS", "git-lfs", []string{"--version"}},

	// Cloud / DevOps
	{"k9s", "k9s", []string{"version"}},
	{"eksctl", "eksctl", []string{"version"}},
	{"terraform-docs", "terraform-docs", []string{"--version"}},

	// Container / Virtualization
	{"Podman", "podman", []string{"--version"}},
	{"Minikube", "minikube", []string{"version"}},

	// Linters / Formatters
	{"eslint_d", "eslint_d", []string{"--version"}},
	{"prettierd", "prettierd", []string{"--version"}},
	{"isort", "isort", []string{"--version"}},

	// Database tools
	{"PostgreSQL", "psql", []string{"--version"}},
	{"MySQL", "mysql", []string{"--version"}},
	{"Redis CLI", "redis-cli", []string{"--version"}},

	{"Crystal", "crystal", []string{"--version"}},
	{"Haskell GHC", "ghc", []string{"--version"}},
	{"OCaml", "ocamlc", []string{"-version"}},
}
