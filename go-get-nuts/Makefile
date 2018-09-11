.PHONY: build run

VERBOSE=1

build:
	GOARCH=wasm GOOS=js go build -o main.wasm

server: build
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
