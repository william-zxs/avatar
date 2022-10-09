package main

import (
	"fmt"
	"github.com/William-ZXS/avatar/internal/config"
	"github.com/William-ZXS/avatar/internal/ssh"
)

// 调用测试
func main() {
	conf := config.ReadConfig()
	for i := 0; i < len(conf.Hosts); i++ {
		host := conf.Hosts[i]
		cli := ssh.NewCli(host.Username, host.Password, host.Addr)
		for _, script := range host.Scripts {
			msg, err := cli.Run(script.Data)
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("msg:", msg)
			}

		}
	}
}
