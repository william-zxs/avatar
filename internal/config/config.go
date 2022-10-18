package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Config struct {
	Hosts []*Host
}

type Host struct {
	Name        string
	Addr        string
	Username    string
	Password    string
	Scripts     []string
	ScriptDatas []*ScriptData
}

type ScriptData struct {
	Name string
	Data string
}

func ReadConfig(file string) *Config {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	viper.Unmarshal(config)
	var cmd strings.Builder
	for i := 0; i < len(config.Hosts); i++ {
		h := config.Hosts[i]
		fmt.Printf("%v", h.Scripts)
		scriptDatas := make([]*ScriptData, 0)
		for i := 0; i < len(h.Scripts); i++ {
			fmt.Println("==script==", h.Scripts[i])
			file := filepath.Join("script", h.Scripts[i]+".sh")
			data, err := os.ReadFile(file)
			checkErr(err)
			cmd.Write(data)
			fmt.Println("==cmd.String()==", cmd.String())
			scriptDatas = append(scriptDatas,
				&ScriptData{
					Name: h.Scripts[i],
					Data: cmd.String(),
				})
			cmd.Reset()
		}
		//for script := range h.Scripts {
		//	fmt.Println("==script==", script)
		//	file := filepath.Join("script", string(script)+".sh")
		//	data, err := os.ReadFile(file)
		//	checkErr(err)
		//	cmd.Write(data)
		//	scriptDatas = append(scriptDatas,
		//		&ScriptData{
		//			Name: string(script),
		//			Data: cmd.String(),
		//		})
		//	cmd.Reset()
		//}
		h.ScriptDatas = scriptDatas
	}
	return config
}
