package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type User struct {
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) GetUser() {
	u := User{
		Username: "yxc",
		Password: "bogzc2002...",
		Hobby:    []string{"basketball", "football"},
	}
	// 返回一个json数据
	c.Data["json"] = u
	c.ServeJSON()

}

func (c *UserController) AddUser() {
	c.TplName = "adduser.html"
}

func (c *UserController) DoAdd() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id必须是int类型")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")
	c.Ctx.WriteString("后端获取到的参数为:" + "id:" + strconv.Itoa(id) + " " + "username:" + username + " " + "password:" + password)

	hobby := c.GetStrings("hobby")
	fmt.Printf("值: %v-- 类型: %T", hobby, hobby)
}

func (c *UserController) EditUser() {
	c.TplName = "edituser.html"
}

func (c *UserController) DoEdit() {
	user := User{}
	if err := c.ParseForm(&user); err != nil {
		c.Ctx.WriteString("post提交失败")
		return
	}
	fmt.Printf("%#v", user)
	c.Ctx.WriteString("解析post数据成功")
}
