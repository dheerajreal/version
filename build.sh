#!/usr/bin/env bash
set -euo pipefail

NAME=version
VERSION=${VERSION:-v0.0.2-dev}
DIST=dist

mkdir -p "$DIST"

export GOFLAGS="-buildvcs=false"
export LC_ALL=C
export TZ=UTC


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
        -trimpath \
        -o "$DIST/${NAME}_${os}_${arch}${ext}" \
        .
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
if command -v sha256sum >/dev/null 2>&1; then
    SUMCMD="sha256sum"
else
    SUMCMD="shasum -a 256"
fi

(
    cd "$DIST"
    $SUMCMD ${NAME}_* > checksums.txt
)

echo "✓ Done"
