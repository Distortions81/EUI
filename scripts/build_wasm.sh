#!/bin/sh
set -e

OUT_DIR="$(dirname "$0")/../web"
mkdir -p "$OUT_DIR"

cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" "$OUT_DIR/"

GOOS=js GOARCH=wasm go build -o "$OUT_DIR/demo.wasm" ../cmd/demo

echo "WASM build output in $OUT_DIR. Open index.html in a browser."
