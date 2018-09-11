.PHONY: build run

VERBOSE=1

build:
	GOARCH=wasm GOOS=js go build -o test.wasm

exec: build
	./go_js_wasm_exec test.wasm

server:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
