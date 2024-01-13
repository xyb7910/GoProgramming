package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	//默认启动方式，其中包含了Logger，Recovery中间件
	//router := gin.Default()

	router := gin.New() //创建一个不包含任何中间键的路由容器

	//全局中间件
	//使用Logger中间键
	router.Use(gin.Logger())
	//使用Recovery中间件
	router.Use(gin.Recovery())

	//路由添加中间件可以添加任意多个
	//router.GET("/benchmark", MyBenchLogger(), benchEndpoint())

	//authorized := router.Group("/", AuthRequested())
	//等价于
	//authorized := router.Group("/")
	//authorized.Use(AuthRequested())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//	authorized.POST("/read", readEndpoint)
	//
	//}

	router.Use(Logger())

	router.GET("/test", test)

	router.Run(":8080")
}

func test(context *gin.Context) {
	example := context.MustGet("example").(string)

	log.Print(example)
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()

		context.Set("example", "12345")

		context.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := context.Writer.Status()
		log.Print(status)
	}
}

func submitEndpoint(context *gin.Context) {

}

func readEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}

//func AuthRequested() gin.HandlerFunc {
//
//}
//
//func benchEndpoint() gin.HandlerFunc {
//
//}
//
//func MyBenchLogger() gin.HandlerFunc {
//
//}
