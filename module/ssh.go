package module

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type Host struct {
	IP       string `yaml:"ip"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
}

func (h *Host) ReadFile(file string) (err error) {
	// var hostKey ssh.PublicKey
	sshConf := &ssh.ClientConfig{
		User: h.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(h.Password),
		},
		// HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	sshClient, err := ssh.Dial("tcp", h.IP+":"+string(h.Port), sshConf)
	if err != nil {
		fmt.Println(err, 1)
		return err
	}
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Println(err, 2)
		return err
	}
	session.Stdout = os.Stdout
	defer session.Close()
	session.Run("date")
	return nil
}
