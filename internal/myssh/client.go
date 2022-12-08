package myssh

import (
	"log"

	"golang.org/x/crypto/ssh"
)

// https://blog.ralch.com/articles/golang-ssh-connection
// https://gist.github.com/iamralch/b7f56afc966a6b6ac2fc
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
	log.Printf("[%s] new ssh connection (%s)", client.RemoteAddr(), client.ClientVersion())

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	log.Printf("[%s] new ssh connection", client.RemoteAddr())

	defer session.Close()
}
