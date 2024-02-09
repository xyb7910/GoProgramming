package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

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
