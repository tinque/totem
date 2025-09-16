#!/bin/bash
set -e

# Get latest git tag, fallback to 'dev' if none
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "dev")
OUTDIR="build"
mkdir -p "$OUTDIR"

platforms=(
  "darwin amd64"
  "darwin arm64"
  "linux amd64"
  "linux arm64"
  "windows amd64"
  "windows arm64"
)

for platform in "${platforms[@]}"; do
  set -- $platform
  GOOS=$1
  GOARCH=$2
  EXT=""
  if [ "$GOOS" = "windows" ]; then EXT=".exe"; fi
  BIN="$OUTDIR/totem-$VERSION-$GOOS-$GOARCH$EXT"
  echo "Building $BIN..."
  env GOOS=$GOOS GOARCH=$GOARCH go build -o "$BIN" main.go
done
echo "All binaries built in $OUTDIR/"