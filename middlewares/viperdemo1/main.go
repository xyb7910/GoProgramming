package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("middlewares/viperdemo1/config/config.yaml") // 指定配置文件路径
	err := viper.ReadInConfig()                                      // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听文件变化
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		// 监听到配置文件变化后，重新读取配置文件
		fmt.Println("The config file is changed:", in.Name)
	})

	r := gin.Default()

	r.GET("/version", func(c *gin.Context) {
		c.String(200, "version: %s", viper.GetString("version"))
	})

	if err := r.Run(fmt.Sprintf(":%d", viper.GetInt("port"))); err != nil {
		panic(err)
	}
}
