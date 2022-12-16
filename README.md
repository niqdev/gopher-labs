# gopher-labs

* [go](https://go.dev/doc) documentation
* [Go by Example](https://gobyexample.com)
* [Standard library](https://pkg.go.dev/std)

TODO
* aws
* kube

## labs

### mydocker

```bash
go run labs.go mydocker --name [run|list]
```

### myssh

```bash
# start server
go run labs.go myssh server

# test
nc 127.0.0.1 2222

# connect with openssh
ssh-keygen -f "/home/ubuntu/.ssh/known_hosts" -R "[localhost]:2222"
ssh -o StrictHostKeyChecking=no foo@localhost -p 2222

# connect with client
go run labs.go myssh client
```

## Development

Setup
```bash
# ubuntu
sudo snap install --classic go

# macos
brew install go

# init project (first time)
go mod init github.com/niqdev/gopher-labs

# install|update dependencies
go mod tidy

# run
go run labs.go
```
