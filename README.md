# gopher-labs

[![ci](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml/badge.svg)](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/niqdev/gopher-labs.svg)](https://pkg.go.dev/github.com/niqdev/gopher-labs)

My Golang laboratory experiments :hourglass_flowing_sand:

## labs

### myconfig

```bash
go run labs.go myconfig
```

### mydocker

```bash
go run labs.go mydocker --name [list|run|attach]
```

### mykube

```bash
# local cluster
minikube start --driver=docker --embed-certs
minikube delete --all

go run labs.go mykube --name [list|exec]

# setup portforward example
# minikube kubectl -- apply -f ./data/install-alpine-xfce-vnc.yaml
go run labs.go mykube --name create

# vncviewer localhost:5900
# http://localhost:6080
go run labs.go mykube --name portforward

# pre-download to solve issue: ErrImagePull (120 seconds timeout)
minikube image load edgelevel/alpine-xfce-vnc:web-0.6.0
```

### mylog

```bash
# zap logging examples
go run labs.go mylog
```

### myschema

```bash
# JSON and Yaml schema validation
go run labs.go myschema
```

### myspinner

```bash
go run labs.go myspinner
```

### myssh

```bash
# start server
go run labs.go myssh server

# test
nc 127.0.0.1 2222

# connect with openssh
ssh-keygen -f ~/.ssh/known_hosts -R "[localhost]:2222"
ssh -o StrictHostKeyChecking=no foo@localhost -p 2222

# connect with client
go run labs.go myssh client
```

### version

```bash
# git version
go run \
  -ldflags="-X github.com/niqdev/gopher-labs/internal.Version=$(git rev-parse HEAD)" \
  labs.go version
```

## Resources

* [go](https://go.dev/doc) documentation
* [Go by Example](https://gobyexample.com)
* [Effective Go](https://github.com/golovers/effective-go)
* [Standard library](https://pkg.go.dev/std)

## Development

Install
```bash
# ubuntu
sudo snap install --classic go

# macos
brew install go
```

Setup
```bash
# init project (first time)
go mod init github.com/niqdev/gopher-labs

# install|update dependencies
go mod tidy

# run
go run labs.go

# list recipes
just

# build
just build
./build/labs
```

Publish
* [Publishing a module](https://go.dev/doc/modules/publishing)
* [How to publish a Go package](https://stackoverflow.com/questions/43716691/how-to-publish-a-go-package)
```bash
git tag vX.Y.Z
git push origin --tags

# refresh index
GOPROXY=proxy.golang.org go list -m github.com/niqdev/gopher-labs@vX.Y.Z

# install
go get github.com/niqdev/gopher-labs
```


TODO
* aws
* http client/server json api
