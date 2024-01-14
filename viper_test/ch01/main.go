package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string `mapstructure:"name"`
	Port        int    `mapstructure:"port"`
}

func main() {
	v := viper.New()

	v.SetConfigFile("viper_test/ch01/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	ServerConfig := ServerConfig{}
	if err := v.Unmarshal(&ServerConfig); err != nil {
		panic(err)
	}
	fmt.Println(ServerConfig)
	fmt.Printf("%v", v.Get("name"))
}
