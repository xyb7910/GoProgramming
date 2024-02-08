package main

import "github.com/gin-gonic/gin"

type Article struct {
	Title   string
	Content string
	Score   int
	Hobby   []string
}

type ArticleList struct {
	articles []Article
}

func main() {
	r := gin.Default()
	// 加载模板
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin/index.html", gin.H{
			"title": "admin index page",
		})
	})

	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(200, "admin/news.html", gin.H{
			"title": "admin news page",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "default/index.html", gin.H{
			"title": "default index page",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		article := Article{
			Title:   "hello",
			Content: "hello world",
			Score:   100,
			Hobby:   []string{"sleep", "eat", "play"},
		}

		articles := []Article{
			{
				Title:   "hello1",
				Content: "hello world",
				Score:   100,
				Hobby:   []string{"sleep", "eat", "play"},
			},
			{
				Title:   "hello2",
				Content: "hello world",
				Score:   100,
				Hobby:   []string{"sleep", "eat", "play"},
			},
			{
				Title:   "hello3",
				Content: "hello world",
				Score:   100,
				Hobby:   []string{"sleep", "eat", "play"},
			},
		}
		c.HTML(200, "default/news.html", gin.H{
			"title":    "default news page",
			"news":     article,
			"articles": articles,
		})
	})
	r.Run()
}
