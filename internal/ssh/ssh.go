package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"time"
)

type Cli struct {
	username string
	password string
	addr     string
	client   *ssh.Client
}

func NewCli(username, password, addr string) Cli {
	return Cli{
		username: username,
		password: password,
		addr:     addr,
	}
}

// Connect 连接远程服务器
func (c *Cli) connect() error {
	config := &ssh.ClientConfig{
		User: c.username,
		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}
	client, err := ssh.Dial("tcp", c.addr, config)
	if err != nil {
		return fmt.Errorf("connect server error: %w", err)
	}
	c.client = client
	return nil
}

// Run 运行命令
func (c Cli) Run(shell string) (string, error) {
	if c.client == nil {
		if err := c.connect(); err != nil {
			return "", err
		}
	}
	session, err := c.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("create new session error: %w", err)
	}
	defer session.Close()
	buf, err := session.CombinedOutput(shell)
	return string(buf), err
}
