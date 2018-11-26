package module

import (
	"bufio"
	"net"

	"github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh"
)

type Host struct {
	IP       string `yaml:"ip"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

func (h *Host) ReadFile(file string, logfile LogFile) (err error) {

	sshConf := &ssh.ClientConfig{
		User: h.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(h.Password),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	sshClient, err := ssh.Dial("tcp", h.IP+":"+string(h.Port), sshConf)
	if err != nil {
		return err
	}
	session, err := sshClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	stdout, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	stdoutread := bufio.NewScanner(stdout)
	stderrread := bufio.NewScanner(stderr)
	go func(logTrans chan []byte) {
		defer close(logTrans)
		for stderrread.Scan() {
			lineerr := stderrread.Text()
			if lineerr != "" {
				logrus.Errorln(err)
			}
			for stdoutread.Scan() {
				logfile.LogTrans <- stdoutread.Bytes()
			}
		}
		return
	}(logfile.LogTrans)
	session.Run("tail -f  " + logfile.Path)

	return nil
}
