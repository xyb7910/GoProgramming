package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Printf("用户名为：%s, 密码：%s", username, password)

	c.Redirect("/", 302)

	//c.Ctx.Redirect(302, "/")
}
