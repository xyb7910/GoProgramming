package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

var config Config

func main() {
	viper.SetConfigFile("middlewares/viperdemo1/config/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()                                      // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 解析配置文件将数据绑定到结构体
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		// 监听到配置文件变化后，重新读取配置文件
		fmt.Println("夭寿啦~配置文件被人修改啦...")
		if err := viper.Unmarshal(config); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	r := gin.Default()

	r.GET("/version", func(c *gin.Context) {
		c.String(200, "version: %s", viper.GetString("version"))
	})

	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}
