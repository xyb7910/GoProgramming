package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	fmt.Println("hello world1")
	start := time.Now().UnixNano()

	c.String(200, "hello world")

	// Nested middleware is executed after this one
	//c.Next()
	c.Abort()

	fmt.Println("hello world2")

	end := time.Now().UnixNano()

	fmt.Printf("middle time: %d\n", end-start)

}

func InitMiddleWareOne(c *gin.Context) {
	fmt.Println("hello world1")
	c.Next()
	fmt.Println("hello world2")
}
func InitMiddleWareTwo(c *gin.Context) {
	fmt.Println("hello world3")
	c.Next()
	fmt.Println("hello world4")
}

func InitMiddleWareThree(c *gin.Context) {
	c.Set("username", "ypb")
	username, _ := c.Get("username")
	fmt.Println(username)

	v, ok := username.(string)
	if ok {
		c.String(200, v)
	} else {
		c.String(200, "获取用户名失败")
	}
}

func main() {
	r := gin.Default()

	// 	全局中间件
	// 使用每一个路由都会调用
	r.Use(InitMiddleWareOne, InitMiddleWareTwo)

	r.GET("/hello", InitMiddleWare, func(c *gin.Context) {
		fmt.Println("hello world3")

		time.Sleep(time.Second * 2)
		c.JSON(200, gin.H{
			"msg": "hello world",
		})
	})

	r.GET("/hello2", InitMiddleWareOne, InitMiddleWareTwo, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world2",
		})
	})

	r.GET("/hello3", InitMiddleWareThree, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello world3",
		})
	})

	r.Run()
}
