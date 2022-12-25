package myssh

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func RunServer() {
	address := MyAddress()
	sshConfig := sshServerConfig()
	listen(address, sshConfig)
}

func sshServerConfig() *ssh.ServerConfig {
	sshConfig := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == myuser && string(pass) == mypassword {
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

	sshConfig.AddHostKey(hostKey)
	return sshConfig
}

func listen(address string, sshConfig *ssh.ServerConfig) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to start ssh server: %v", err)
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
