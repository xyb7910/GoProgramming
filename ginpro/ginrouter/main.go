package main

import (
	"github.com/gin-gonic/gin"

	"LearingGo/ginpro/ginrouter/router"
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
