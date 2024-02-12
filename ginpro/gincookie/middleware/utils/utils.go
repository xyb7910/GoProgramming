package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (con DefaultController) SET(c *gin.Context) {
	c.SetCookie("username", "李四", 3600, "/", "localhost", false, true)
	c.SetCookie("age", "18", 10, "/", "localhost", false, true)
	c.HTML(http.StatusOK, "templates/index.html", gin.H{
		"msg": "我是一个msg",
		"t":   1629788010,
	})
}

func (con DefaultController) GET(c *gin.Context) {
	//获取cookie
	username, _ := c.Cookie("username")
	age, _ := c.Cookie("age")
	c.String(http.StatusOK, "用户--cookie.username="+username)
	c.String(http.StatusOK, "用户--cookie.age="+age)
}

func (con DefaultController) LOGIN1(c *gin.Context) {
	c.SetCookie("username", "李四", 3600, "/", "*.gin.com", false, true)
	c.String(200, "login success")
}

func (con DefaultController) GET2(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.String(http.StatusOK, "用户--cookie.username="+username)
}

// DeleteCookie 删除cookie
func (con DefaultController) DeleteCookie(c *gin.Context) {
	c.SetCookie("username", "李四", -1, "/", "127.0.0.1", false, true)
}
