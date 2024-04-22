package main

import (
	"ginpro/ginmodel/models"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	// 注册模板函数必须在加载模板上面
	r.SetFuncMap(template.FuncMap{
		"print": models.Print,
	})

	r.LoadHTMLGlob("templates/*")
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(200, "ginmodel/deafult.html", gin.H{
			"ID":      "1",
			"message": "Hello World",
			"func":    models.Print,
		})
	})
	r.Run()
}
