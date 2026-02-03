#!/usr/bin/env bash
set -euo pipefail

NAME=version
VERSION=0.0.1-dev
DIST=dist

mkdir -p "$DIST"

build() {
    local os=$1
    local arch=$2
    local ext=$3
    
    echo "→ Building $os/$arch"
    
    CGO_ENABLED=0 \
    GOOS=$os \
    GOARCH=$arch \
    go build \
    -ldflags="-s -w -X main.version=$VERSION" \
    -o "$DIST/${NAME}_${VERSION}_${os}_${arch}${ext}"
}

# Linux
build linux amd64 ""
build linux arm64 ""

# macOS
build darwin amd64 ""
build darwin arm64 ""

# Windows
build windows amd64 ".exe"

# Checksums
(
    cd "$DIST"
    sha256sum * > checksums.txt
)

echo "✓ Done"
