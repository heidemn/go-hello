all: go-wasm.exe wasm_exec.js

go-wasm.exe:
	GOOS=js GOARCH=wasm go build -o go-wasm.wasm

wasm_exec.js:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .
