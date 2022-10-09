package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Hosts []*Host
}

type Host struct {
	Name        string
	Addr        string
	Username    string
	Password    string
	ScriptNames []string
	Scripts     []*Script
}

type Script struct {
	Kind string
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

	for i := 0; i < len(config.Hosts); i++ {
		host := config.Hosts[i]
		for _, scriptName := range host.ScriptNames {
			viper.SetConfigName(scriptName)
			viper.SetConfigType("yaml")
			viper.AddConfigPath("script")
			err = viper.ReadInConfig()

			script := &Script{}
			err = viper.Unmarshal(script)
			host.Scripts = append(host.Scripts, script)
		}
	}
	return config
}
