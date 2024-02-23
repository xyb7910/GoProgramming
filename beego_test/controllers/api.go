package controllers

import beego "github.com/beego/beego/v2/server/web"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) GetApi() {
	// 获取动态路由
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api 接口 ---" + id)
}

func (c *ApiController) GetCms() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("CMS详情---" + id)
}
