package main

import (
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	Decription string `json:"description"`
}

func main() {

	// 建立一个 gin 服务器
	r := gin.Default()
	// 加载模板
	r.LoadHTMLGlob("templates/*")

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	// map
	r.GET("/JSON1", func(c *gin.Context) {
		c.JSON(200, map[string]interface{}{
			"name": "Jack",
			"age":  20,
			"sex":  "male",
		})
	})

	// gin.H
	r.GET("/JSON2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": "Alice",
			"age":  21,
			"sex":  "male",
		})
	})

	// struct
	r.GET("/JSON3", func(c *gin.Context) {

		a := &Article{
			Title:      "Hello",
			Content:    "This is a content",
			Decription: "This is a decription",
		}
		c.JSON(200, a)
	})

	//  JSONP
	// http://127.0.0.1:8080/JSONP?callback=XXX  XXX为回调函数名
	// 主要是用于JSONP的跨域请求
	r.GET("/JSONP", func(c *gin.Context) {

		a := &Article{
			Title:      "Hello",
			Content:    "This is a content",
			Decription: "This is a decription",
		}
		c.JSONP(200, a)
	})

	// XML
	r.GET("/XML", func(c *gin.Context) {
		c.XML(200, map[string]interface{}{
			"name": "Jack",
			"age":  20,
			"sex":  "male",
		})
	})

	// XML1
	r.GET("/XML1", func(c *gin.Context) {
		c.XML(200, gin.H{
			"name": "Alice",
			"age":  21,
			"sex":  "male",
		})
	})

	// XML2
	r.GET("/XML2", func(c *gin.Context) {
		a := &Article{
			Title:      "Hello",
			Content:    "This is a content",
			Decription: "This is a decription",
		}
		c.XML(200, a)
	})

	// HTML
	// use html must load index.html in the same folder
	// the LoadHTMLGlob() method loads all HTML files
	// 加载模板,注意templates文件夹下必须有index.html文件
	r.GET("/HTML", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Hello World!",
		})
	})

	// 运行服务器
	if err := r.Run(); err != nil {
		panic(err)
	}
}
