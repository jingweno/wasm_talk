.PHONY: default build exec server

VERBOSE=1

default: server

build:
	GOARCH=wasm GOOS=js go build -o test.wasm

exec: build
	./go_js_wasm_exec test.wasm

server:
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
