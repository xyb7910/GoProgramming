package main

import (
	"LearingGo/gormpro/link_database/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/add", controller.UserController{}.Add, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "add success",
		})
	})

	r.GET("/get", controller.UserController{}.Get, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get success",
		})
	})

	r.GET("/getbyname", controller.UserController{}.GetByCondition, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get success",
		})
	})

	r.GET("/update", controller.UserController{}.Update, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update success",
		})
	})

	r.GET("/updatebycondition", controller.UserController{}.UpdateByCondition, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "update success",
		})
	})

	r.GET("/delete", controller.UserController{}.Delete, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete success",
		})
	})

	r.GET("/deleteall", controller.UserController{}.DeleteAll, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete success",
		})
	})

	r.GET("addarticle", controller.ArticleController{}.Add, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "add article success",
		})
	})
	r.GET("getarticle", controller.ArticleController{}.GetOne, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get article success",
		})
	})

	r.GET("getsum", controller.ArticleController{}.GetSum, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get article success",
		})
	})

	r.GET("getscan", controller.ArticleController{}.GetScan, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get article success",
		})
	})
	r.Run()
}
