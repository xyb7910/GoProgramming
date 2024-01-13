package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置模板
	r.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	r.StaticFS("/static", http.Dir("./static"))
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/posts/index", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("gets/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/login.tmpl", gin.H{
			"title": "gets/login",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()
}
