BUILD_PATH := "./build"
BIN_NAME := "labs"

default:
  @just --list

format:
  go fmt ./...

test:
	go test ./...

build: format test
  rm -frv {{BUILD_PATH}}
  go build \
    -ldflags="-X github.com/niqdev/gopher-labs/internal.Version=$(git rev-parse HEAD)" \
    -o {{BUILD_PATH}}/{{BIN_NAME}} labs.go
