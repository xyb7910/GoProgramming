package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

type Article struct {
	Title   string
	Content string
	Score   int
	Hobby   []string
}

type ArticleList struct {
	articles []Article
}

// UnixToTime 时间转换
func UnixToTime(timestamp int) string {
	fmt.Println(int64(timestamp))
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func PrintToAdd(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
}

func main() {
	r := gin.Default()
	// 自定义函数, 必须放在加载模板前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"PrintToAdd": PrintToAdd,
	})
	// 加载模板
	r.LoadHTMLGlob("templates/**/*")
	//配置静态资源, 浏览器访问链接为 http://localhost:8080/static/css/base.css
	r.Static("/static", "./static")

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
			"title":     "default index page",
			"content":   "hello world",
			"time":      time.Now(),
			"timestamp": 1707388414,
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
