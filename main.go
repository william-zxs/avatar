package main

import (
	"flag"
	"fmt"
	"github.com/William-ZXS/avatar/internal/config"
	"github.com/William-ZXS/avatar/internal/ssh"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	confFile := flag.String("f", "conf.yaml", "请指定配置文件")
	flag.Parse()
	//加载配置、脚本
	conf := config.ReadConfig(*confFile)

	//执行脚本
	for i := 0; i < len(conf.Hosts); i++ {
		host := conf.Hosts[i]
		cli := ssh.NewCli(host.Username, host.Password, host.Addr)
		for _, scriptData := range host.ScriptDatas {
			msg, err := cli.Run(scriptData.Data)
			if err != nil {
				fmt.Println("err:", err)
			} else {
				fmt.Println("msg:", msg)
			}
		}
	}

}
