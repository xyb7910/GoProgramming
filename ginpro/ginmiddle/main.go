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

func TestCtxNext(c *gin.Context) {
	fmt.Println("1-执行中间件")
	start := time.Now().UnixNano()
	// 执行下一个中间件
	c.Next()
	fmt.Println("3-程序执行完成，计算时间")
	// 计算时间
	end := time.Now().UnixNano()
	fmt.Printf("middle time: %d\n", end-start)
}

// MiddleWareSetValue 设置值
func MiddleWareSetValue(c *gin.Context) {
	// 可以通过 ctx.Set 在请求上下文中设置值，后续的处理函数能够取到该值
	c.Set("username", "ypb")
	c.String(200, "中间件设置值")
}

type Controller struct{}

// ControllerGetValue 获取值
func (c Controller) ControllerGetValue(ctx *gin.Context) {
	username, _ := ctx.Get("username")
	fmt.Println(username)
	ctx.String(200, "获取用户名成功")
}

func main() {
	r := gin.Default()

	// 全局中间件
	// 使用每一个路由都会调用
	//r.Use(InitMiddleWareOne, InitMiddleWareTwo)

	r.GET("/getvalue", MiddleWareSetValue, Controller{}.ControllerGetValue)

	r.GET("/", TestCtxNext, func(c *gin.Context) {
		fmt.Println("2-执行路由")
		time.Sleep(time.Second)
		c.String(200, "hello world")
	})

	r.GET("/hello", InitMiddleWare, func(c *gin.Context) {
		fmt.Println("hello world3")
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
