#!/bin/sh
set -e

OUT_DIR="$(dirname "$0")/../web"
mkdir -p "$OUT_DIR"

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$OUT_DIR/"

# Build with size optimization flags. -s and -w strip debug info
# and -trimpath removes filesystem paths for a smaller binary.
GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o "$OUT_DIR/demo.wasm" ../cmd/demo

# Compress the wasm binary with Brotli for smaller downloads
brotli -f "$OUT_DIR/demo.wasm" -o "$OUT_DIR/demo.wasm.br"

echo "WASM build output in $OUT_DIR. Open index.html in a browser."
