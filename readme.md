# 💻 Version Checker CLI


A fast, simple, and oh-so-helpful CLI to check versions of your favorite programming tools, package managers, dev toys, and cloud goodies.


> “Because who has time to --version everything manually? or is it -v ?”

Bonus: get path of the cli in output. No more wondering `where` 😕 it came from.

## 🚀 Features

Detects versions for languages, package managers, dev tools, and cloud CLIs.

Prints a single tool’s version or all tools at once (because patience is overrated).

Supports 80+ tools (it's basically the version bro of your system 😏).

Outputs as a pretty table or JSON for your inner data nerd.

Handles missing tools gracefully — no crashes, just sass-free "Not Found" messages.

Built in Go for speed and elegance


## 🛠 Installation

Using Go, like a pro
```
go install github.com/dheerajreal/version/cmd/version@latest
```


Make sure `$GOPATH/bin` or `$GOBIN` is in your `PATH` — your shell deserves it. 💃

## 🧩 Usage

```
# Show the help message (aka cheat sheet)
version

# Check the version of the CLI itself (because we’re meta)
version --version

# Check a single tool
version python

# Show all tools at once (you monster)
version --all

# Output in JSON format (for the nerds)
version --all --json
```


## 💅 Examples

Check Python version:

```
$ version python
Python          3.11.4               /usr/bin/python3
```

Check all tools in JSON format (nerd alert 🚨):

```
$ version --all --json
[
  {
    "name": "Python",
    "version": "3.11.4",
    "path": "/usr/bin/python3"
  },
  {
    "name": "Node.js",
    "version": "20.5.1",
    "path": "/usr/bin/node"
  },
  ...
]
```


## 💎 Supported Tools

Languages: Python, Node.js, Go, Rust, Java, Ruby, PHP, Swift, Kotlin, Dart, Perl, Lua, Elixir, R, Julia, C/C++, Clang, Zig, Crystal, Haskell, OCaml

Package Managers: npm, yarn, pnpm, bun, pip, pipx, uv, cargo, gem, composer, Homebrew, luarocks, opam, pub

Dev Tools: Git, Git LFS, Docker, kubectl, Helm, Terraform, Ansible, Vagrant, Make, CMake, Gradle, Maven, ESLint, Prettier, Black, Ruff, SQLite, jq, curl

Cloud: AWS CLI, GCloud, Azure CLI, Fly.io, Vercel, Netlify, Railway, Render, Cloudflare Wrangler, Firebase

Containers & Virtualization: Podman, Minikube

Linters & Formatters: eslint_d, prettierd, isort

Databases: PostgreSQL, MySQL, Redis CLI

Basically, if it has a version — we check it. 💁‍♂️

## 🤝 Contributing

Think you can make it faster or more fabulous?

- Add new tools 🔧

- Make the CLI even cooler ✨

- Pull requests and star-gazing welcome ⭐

## 🏷 License

MIT License © 2026 Dheeraj