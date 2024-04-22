package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个 gin.Engine
	r := gin.Default()

	// 注册路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	r.GET("/hi", func(c *gin.Context) {
		c.String(200, "hi world")
	})

	r.GET("/test/:age", func(c *gin.Context) {
		age := c.Query("age")
		c.String(200, "I am %s years old", age)
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

	r.GET("jsonp", func(c *gin.Context) {
		data := gin.H{
			"message": "hello world",
		}
		c.JSONP(200, data)
	})

	// 数据返回为XML形式
	r.GET("/getXml", func(c *gin.Context) {
		// 方式一自己拼接xml
		c.XML(200, gin.H{"name": "john", "age": 20})
	})

	r.GET("/getXml2", func(c *gin.Context) {
		// 方式二使用结构体
		type Message struct {
			Name    string
			Message string
			Age     int
		}
		var msg Message
		msg.Name = "john"
		msg.Message = "hello world"
		msg.Age = 20
		c.XML(200, msg)
	})
	// run 默认的服务地址为 127.0.0.1:8080
	r.Run("127.0.0.1:8080")
}
