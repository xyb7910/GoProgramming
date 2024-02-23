package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) GetArticle() {
	c.Ctx.WriteString("这是获取文章的页面")
	id, err := c.GetInt("id")
	if err == nil {
		fmt.Printf("值: %d, 类型:%T", id, id)
		c.Ctx.WriteString("修改了新闻")
	} else {
		c.Ctx.WriteString("闯入参数有误")
	}
}

func (c *ArticleController) EditArticle() {
	c.Ctx.WriteString("这是编辑文章的页面")
}
