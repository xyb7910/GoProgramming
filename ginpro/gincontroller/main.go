package main

import (
	"LearingGo/ginpro/gincontroller/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化路由
	r := gin.Default()

	// 注册路由
	router.AdminRoutersInit(r)
	router.ApiRoutersInit(r)
	router.DefaultRoutersInit(r)

	// 运行服务器
	r.Run(":8080")
}
