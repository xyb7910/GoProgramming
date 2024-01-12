package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/welcome", welcome) //从get请求中获取参数
	router.POST("/from_post", formPost)
	router.POST("/post", getPost)
	router.Run(":8080")
}

func getPost(context *gin.Context) {
	id := context.Query("id")
	page := context.DefaultQuery("page", "0")
	name := context.PostForm("name")
	message := context.PostForm("message")
	context.JSON(200, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

func formPost(context *gin.Context) {
	message := context.PostForm("message")
	nick := context.DefaultPostForm("nick", "anonymous")
	context.JSON(200, gin.H{
		"message": message,
		"nick":    nick,
	})
}

func welcome(context *gin.Context) {
	firstName := context.DefaultQuery("firstname", "???")
	lastName := context.Query("lastname")
	context.JSON(200, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})
}
