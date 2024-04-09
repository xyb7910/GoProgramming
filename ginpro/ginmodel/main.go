package main

import (
	"ginpro/ginmodel/models"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.SetFuncMap(template.FuncMap{
		"print": models.Print,
	})

	r.GET("/hello", func(c *gin.Context) {
		c.HTML(200, "ginmodel/deafult.html", gin.H{
			"ID":      "1",
			"message": "Hello World",
			"func":    models.Print,
		})
	})
	r.Run()
}
