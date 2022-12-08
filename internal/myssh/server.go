package myssh

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"golang.org/x/crypto/ssh"
)

// https://scalingo.com/blog/writing-a-replacement-to-openssh-using-go-12
// https://blog.gopheracademy.com/advent-2015/ssh-server-in-go
// https://gist.github.com/protosam/53cf7970e17e06135f1622fa9955415f
// https://github.com/ContainerSSH/MiniContainerSSH/blob/master/main.go
// https://github.com/gogs/gogs/blob/main/internal/ssh/ssh.go
func RunServer() {
	const (
		host = "0.0.0.0"
		port = 2222
	)
	address := net.JoinHostPort(host, strconv.Itoa(port))
	sshConfig := sshServerConfig()
	listen(address, sshConfig)
}

func sshServerConfig() *ssh.ServerConfig {
	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == "foo" && string(pass) == "bar" {
				// TODO return metadata to the client
				return &ssh.Permissions{Extensions: map[string]string{"user-id": c.User()}}, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
		//NoClientAuth: true,
	}

	// TODO read from file
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln(err)
	}
	hostKey, err := ssh.NewSignerFromKey(key)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("new private host key generated")

	config.AddHostKey(hostKey)
	return config
}

func listen(address string, sshConfig *ssh.ServerConfig) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to start SSH server: %v", err)
	}
	log.Printf("listening on %s", address)

	for {
		tcpConnection, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting incoming connection: %v", err)
			// do not stop the server
			continue
		}

		go func() {
			log.Printf("[%s] new ssh handshake", tcpConnection.RemoteAddr())
			// TODO channels, requests
			sshConnection, _, _, err := ssh.NewServerConn(tcpConnection, sshConfig)
			if err != nil {
				if err == io.EOF {
					log.Printf("[%s] handshake terminated: %v", tcpConnection.RemoteAddr(), err)
				} else {
					log.Printf("[%s] handshake error: %v", tcpConnection.RemoteAddr(), err)
				}
				return
			}

			log.Printf("[%s] new ssh connection (%s)", sshConnection.RemoteAddr(), sshConnection.ClientVersion())

			sshConnection.Close()
		}()
	}
}
