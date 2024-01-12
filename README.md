# gopher-labs

[![ci](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml/badge.svg)](https://github.com/niqdev/gopher-labs/actions/workflows/ci.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/niqdev/gopher-labs.svg)](https://pkg.go.dev/github.com/niqdev/gopher-labs)

My Golang laboratory experiments :hourglass_flowing_sand:

## labs

### myargo

```bash
# lists argo-cd applications
go run labs.go myargo list

# submits argo workflow
go run labs.go myargo submit
```

### myaws

* [AWS SDK for Go](https://aws.github.io/aws-sdk-go-v2/docs) documentation
* [AWS SDK for Go V2 code examples for Amazon SQS](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/sqs)

```bash
make local-up

curl http://localhost:4566/health | jq

# verify hooks
curl -s localhost:4566/_localstack/init | jq .

# create (see entrypoint folder)
docker exec -it go-dev aws --endpoint-url=http://localstack:4566 sqs create-queue \
  --queue-name go-sqs-example

# list
docker exec -it go-dev aws --endpoint-url=http://localstack:4566 sqs list-queues
docker exec -it go-localstack awslocal sqs list-queues

# produce
docker exec -it go-dev aws --endpoint-url=http://localstack:4566 sqs send-message \
  --queue-url http://localstack:4566/000000000000/go-sqs-example \
  --message-body "hello"

# consume
docker exec -it go-dev aws --endpoint-url=http://localstack:4566 sqs receive-message \
  --queue-url http://localstack:4566/000000000000/go-sqs-example

go run labs.go myaws --name sqs-write
go run labs.go myaws --name sqs-read
```

### myconcurrency

```bash
# uses goroutine and channels
go run labs.go myconcurrency
```

### myconfig

```bash
# prints config examples
go run labs.go myconfig
```

### mydocker

```bash
go run labs.go mydocker --name [list|run|attach]
```

### myhttp

* [Top Go Web Frameworks](https://github.com/mingrammer/go-web-framework-stars)

```bash
# client
docker run -p 8080:80 kennethreitz/httpbin
go run labs.go myhttp --name client

# server
go run labs.go myhttp --name server
curl -Lv http://localhost:3333/home
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

# create job and tail logs
go run labs.go mykube --name job

# copy README.md to/from pod
go run labs.go mykube --name copy-to
go run labs.go mykube --name copy-from
```

### mylog

```bash
# zap logging examples
go run labs.go mylog
```

### myschema

```bash
# json and yaml schema validation
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
* [Awesome Go](https://github.com/avelino/awesome-go)

## Development

Install
```bash
# ubuntu
sudo snap install --classic go

# macos
brew install go

# update version
go mod edit -go 1.21
```

Setup
```bash
# init project (first time)
go mod init github.com/niqdev/gopher-labs

# update dependencies
go mod tidy

# install dependencies (examples)
go get github.com/argoproj/argo-cd/v2
go get github.com/argoproj/argo-workflows/v3
go get -u github.com/hashicorp/go-retryablehttp
go get github.com/onsi/ginkgo/v2
go get github.com/onsi/gomega/...
go get github.com/aws/aws-sdk-go-v2
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/sqs
go mod vendor

# run
go run labs.go

# list recipes
just

# build
just build
./build/labs

# test
just test
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
