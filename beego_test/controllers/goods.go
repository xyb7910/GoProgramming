package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type Good struct {
	Title   string `form:"title" json:"title" xml:"title"`
	Content string `form:"content" json:"content" xml:"content"`
}

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.TplName = "goods.html"
}

func (c *GoodsController) DoEdit() {
	c.Ctx.WriteString("执行修改商品")
}

func (c *GoodsController) DoDelete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("参数错误")
	}
	c.Ctx.WriteString("执行删除操作--" + strconv.Itoa(id))
}
