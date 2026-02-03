#!/bin/bash
set -e

# Output folder
mkdir -p dist

# Linux x86_64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/version-linux-amd64 main.go
# Linux ARM64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o dist/version-linux-arm64 main.go
# macOS x86_64
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o dist/version-macos-amd64 main.go
# macOS ARM64 (M1/M2)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o dist/version-macos-arm64 main.go
# Windows x86_64
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/version-windows-amd64.exe main.go
