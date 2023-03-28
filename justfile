BUILD_PATH := "./build"
BIN_NAME := "labs"

# https://stackoverflow.com/questions/61515186/when-using-cgo-enabled-is-must-and-what-happens
GO_BUILD_ENV := "CGO_ENABLED=0 GOOS=linux GOARCH=amd64"
GO_FILES := "./..."

default:
  @just --list

format:
  go fmt {{GO_FILES}}

vet:
  go vet {{GO_FILES}}

test:
  go test {{GO_FILES}} -cover

build $VERSION_COMMIT="$(git rev-parse HEAD)": format test
  rm -frv {{BUILD_PATH}}
  go build \
    -ldflags="-X github.com/niqdev/gopher-labs/internal.Version={{VERSION_COMMIT}}" \
    -o {{BUILD_PATH}}/{{BIN_NAME}} labs.go
