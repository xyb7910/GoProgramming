package main

import (
	"github.com/gin-gonic/gin"
	"path"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "template/index.html", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {
		username := c.PostForm("username")

		file, err := c.FormFile("face")
		dst := path.Join("./static/upload", file.Filename)
		if err == nil {
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"dst":      dst,
		})
	})

	r.GET("/any", func(c *gin.Context) {
		c.HTML(200, "templates/default.html", gin.H{})
	})

	r.POST("/uploadfiles", func(c *gin.Context) {
		username := c.PostForm("username")

		file1, err := c.FormFile("face1")
		dst := path.Join("./static/upload", file1.Filename)
		if err == nil {
			c.SaveUploadedFile(file1, dst)
		}
		file2, err := c.FormFile("face2")
		dst1 := path.Join("./static/upload", file2.Filename)
		if err == nil {
			c.SaveUploadedFile(file1, dst1)
		}
		c.JSON(200, gin.H{
			"success":  true,
			"username": username,
			"dst":      dst,
			"dst1":     dst1,
		})
	})
	r.Run()
}
