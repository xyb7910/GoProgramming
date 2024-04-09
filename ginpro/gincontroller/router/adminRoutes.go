package router

import (
	"ginpro/gincontroller/controller/admin"
	"github.com/gin-gonic/gin"
)

// AdminRoutersInit  管理后台路由
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.GetUserList)
		adminRouters.GET("/article", admin.ArticleController{}.GetArticle)
	}
}
