package myssh

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func RunClient() {
	address := MyAddress()
	sshConfig := sshClientConfig()
	connect(address, sshConfig)
}

func sshClientConfig() *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: myuser,
		Auth: []ssh.AuthMethod{
			ssh.Password(mypassword),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return sshConfig
}

func connect(address string, sshConfig *ssh.ClientConfig) {
	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	log.Printf("[%s] ssh connection estabilished (%s)", client.RemoteAddr(), client.ClientVersion())

	defer session.Close()
}
