.PHONY: default build exec server

VERBOSE=1

default: serve

generate:
	protoc -I. ./proto/web.proto \
		--go_out=plugins=grpc:$$GOPATH/src
	go generate -x ./frontend/

serve:
	go run main.go
