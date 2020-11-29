#!/bin/bash
GOOS=js GOARCH=wasm go build -o go-wasm.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
