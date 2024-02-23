package routers

import (
	"beego_test/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 动态路由
	beego.Router("/api/:id", &controllers.ApiController{}, "get:GetApi")
	// 正则表达式
	beego.Router("/cms_:id([0-9]+).html", &controllers.ApiController{}, "get:GetCms")

	beego.Router("/register", &controllers.RegisterController{}, "get:Get")
	beego.Router("/doregister", &controllers.RegisterController{}, "post:DoRegister")
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/dologin", &controllers.LoginController{}, "post:DoLogin")

	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/doedit", &controllers.GoodsController{}, "put:DoEdit")
	beego.Router("/goods/delete:id", &controllers.GoodsController{}, "delete:DoDelete")

	beego.Router("/article/get:id", &controllers.ArticleController{}, "get:GetArticle")
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")

	beego.Router("/user", &controllers.UserController{}, "get:GetUser")
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/get", &controllers.UserController{}, "get:GetUser")
	beego.Router("/user/doadd", &controllers.UserController{}, "post:DoAdd")
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditUser")
	beego.Router("/user/doedit", &controllers.UserController{}, "post:DoEdit")
}
