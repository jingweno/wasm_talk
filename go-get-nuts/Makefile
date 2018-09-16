.PHONY: default build server

VERBOSE=1

default: server

build:
	GOARCH=wasm GOOS=js go build -o main.wasm

server: build
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))'
