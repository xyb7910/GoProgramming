package main

import (
	"fmt"
	"geektime_test_lesson/internal/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	// 实现跨域问题
	//server.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//	AllowHeaders:     []string{"Content-Type"},
	//	AllowOriginFunc: func(origin string) bool {
	//		if strings.HasPrefix(origin, "http://localhost") {
	//			return true
	//		}
	//		return strings.Contains(origin, "your_company.com")
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	// 实现 user 相关路由的注册
	user := web.UserHandler{}
	user.RegisterRoute(server)

	server.GET("/ping", func(c *gin.Context) {
		// 输出请求路径
		fmt.Println(c.Request.URL.Path)
		c.String(http.StatusOK, "Hello World")
	})

	// 静态路由匹配
	server.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "这是 hello 请求的路由返回结果")
	})

	// 参数路由
	server.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "这是你传过来的名字: %s", name)
	})

	// 通配符路由
	server.GET("/views/*.html", func(c *gin.Context) {
		path := c.Param(".html")
		c.String(http.StatusOK, "这是你传过来的路径: %s", path)
	})

	// 查询参数
	server.GET("/order", func(c *gin.Context) {
		// 获取查询参数
		id := c.Query("id")
		c.String(http.StatusOK, "这是你传过来的订单id: %s", id)
	})
	_ = server.Run(":8080")
}
