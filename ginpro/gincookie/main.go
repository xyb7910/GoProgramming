package main

import (
	"github.com/gin-gonic/gin"

	"LearingGo/ginpro/gincookie/middleware/utils"
)

func main() {
	r := gin.Default()

	r.GET("/set", utils.DefaultController{}.SET, func(c *gin.Context) {
		c.String(200, "set finish")
	})

	r.GET("/get", utils.DefaultController{}.GET, func(c *gin.Context) {
		c.String(200, "get finish")
	})

	r.GET("/delete", utils.DefaultController{}.DeleteCookie, func(c *gin.Context) {
		c.String(200, "delete finish")
	})

	/*
		设置 二级路由
	*/

	r.GET("/login", utils.DefaultController{}.LOGIN1, func(c *gin.Context) {
		c.String(200, "login finish")
	})

	r.GET("/get1.gin.com", utils.DefaultController{}.GET2, func(c *gin.Context) {
		c.String(200, "get finish")
	})
	r.GET("/get2.gin.com", utils.DefaultController{}.GET2, func(c *gin.Context) {
		c.String(200, "get finish")
	})

	r.Run()
}
