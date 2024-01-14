package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type ServerConfig struct {
	Name      string      `mapstructure:"name"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
}

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func main() {
	//实现生产环境上的分离
	data := GetEnvInfo("Debug")
	var configFileName string
	configFileNamePrefix := "config"
	if data == "true" {
		configFileName = fmt.Sprintf("viper_test/%s-debug.yaml", configFileNamePrefix)
	} else {
		configFileName = fmt.Sprintf("viper_test/%s-pro.yaml", configFileNamePrefix)
	}

	serverConfig := ServerConfig{}

	fmt.Println(data)

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	//动态监听变化
	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
			_ = v.ReadInConfig() // 读取配置数据
			_ = v.Unmarshal(&serverConfig)
			fmt.Println(serverConfig)
		})
	}()

	time.Sleep(time.Second * 3000)

}
