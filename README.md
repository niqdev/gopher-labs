# gopher-labs

[![ci](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml/badge.svg)](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml)

* [go](https://go.dev/doc) documentation
* [Go by Example](https://gobyexample.com)
* [Standard library](https://pkg.go.dev/std)

## labs

### mydocker

```bash
go run labs.go mydocker --name [run|list|attach]
```

<!--
### mykube

```bash
# local cluster
minikube start --driver=docker --embed-certs
minikube delete --all

go run labs.go mykube --name [create]
```
--->

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

# build
just
```

TODO
* kube
* aws
* config
* client/server http json api
* expose pkg
* add docs
* vet/fmt/lint action
* tests
