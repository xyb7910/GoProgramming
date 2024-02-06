package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 gin.Engine
	r := gin.Default()

	// r.

	// 注册路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.GET("/hi", func(c *gin.Context) {
		c.String(200, "hi world")
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello %s", name)
	})

	r.POST("/post", func(c *gin.Context) {
		c.String(200, "post")
	})

	r.PUT("/put", func(c *gin.Context) {
		c.String(200, "put 主要用于更新数据")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "delete 主要用于删除数据")
	})

	// run 默认的服务地址为 127.0.0.1:8080
	r.Run("127.0.0.1:8080")
}
