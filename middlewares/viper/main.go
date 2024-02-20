package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	r := gin.Default()
	r.Run(":8080")
	// 建立默认值

	//viper.SetDefault("Content", "Hello World")
	//fmt.Println(viper.GetString("Content"))

	//viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config") // 设置配置文件名称
	viper.SetConfigType("yaml")   // 设置配置文件类型
	// middlewares/viper/config
	viper.AddConfigPath("middlewares/viper/config") // 设置配置文件路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件不存在
			fmt.Println("The config file is not found")
		} else {
			// 配置文件找到，但读取失败
			fmt.Println("The config file is found but read failed")
		}
	}
	// 配置文件读取成功
	// fmt.Println("The config file is read successfully")
	viper.WatchConfig() // 监听配置文件变化
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 监听到配置文件变化后，重新读取配置文件
		fmt.Println("The config file is changed:", in.Name)
	})
}
