package controller

import (
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

// SetSession set session
func (con DefaultController) SetSession(c *gin.Context) {
	// set session
	session := sessions.Default(c)

	//设置过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, //6hrs
	})

	session.Set("name", "lemon")
	session.Save()

	c.String(200, "set session name lemon")
}

// GetSession get session
func (con DefaultController) GetSession(c *gin.Context) {
	// 初始化session
	session := sessions.Default(c)
	name := session.Get("name")
	c.String(200, "get session name %s", name)
}
