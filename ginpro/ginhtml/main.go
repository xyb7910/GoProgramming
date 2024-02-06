package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 加载模板
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "index page",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(200, "news.html", gin.H{
			"title": "news page",
		})
	})

	r.Run()
}
