BUILD_PATH := "./build"
BIN_NAME := "labs"

default: build

format:
  go fmt ./...

build: format
  rm -frv {{BUILD_PATH}}
  go build \
    -ldflags="-X github.com/niqdev/gopher-labs/internal.Version=$(git rev-parse HEAD)" \
    -o {{BUILD_PATH}}/{{BIN_NAME}} labs.go
