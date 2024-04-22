package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo1 struct {
	Name string `form:"name" json:"name" binding:"required"`
	Age  int    `form:"age" json:"age" binding:"required"`
}

type UserInfo2 struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Article struct {
	Title   string `xml:"title" binding:"required"`
	Content string `xml:"content" binding:"required"`
}

func main() {
	r := gin.Default()

	// 加载模板
	r.LoadHTMLGlob("templates/*")

	//user := struct {
	//	Name   string
	//	Gender string
	//	Age    int
	//}{
	//	Name:   "john",
	//	Gender: "male",
	//	Age:    20,
	//}
	//r.GET("/getuser", func(c *gin.Context) {
	//	c.HTML(200, "templates/index.html", gin.H{
	//		"user": user,
	//	})
	//})

	// GET /?name=john&age=20&sex=male&address=shanghai
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		sex := c.Query("sex")
		address := c.DefaultQuery("address", "shanghai")
		c.JSON(200, gin.H{
			"name":    name,
			"age":     age,
			"sex":     sex,
			"address": address,
		})
	})

	// 获取POST请求的数据
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "templates/index.html", gin.H{})
	})

	r.POST("/doAddUser1", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	// 将数据绑定到UserInfo结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo1{}
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("user: %+v\n", user)
			c.JSON(200, user)
		}
	})

	r.POST("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo2{}
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("user: %+v\n", user)
			c.JSON(200, user)
		}
	})

	// 获取POST xml请求的数据
	r.POST("/getUserXml", func(c *gin.Context) {
		article := &Article{}

		xmlSliceData, _ := c.GetRawData()

		if err := xml.Unmarshal(xmlSliceData, &article); err == nil {
			c.JSON(200, article)
		} else {
			c.JSON(400, gin.H{"err": err.Error()})
		}
	})

	r.Run()
}
