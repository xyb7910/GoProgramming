package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) DoRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	repassword := c.GetString("repassword")

	fmt.Printf("用户名为：%s, 密码：%s, 确认密码为：%s", username, password, repassword)

	c.Ctx.Redirect(302, "/login")
}
