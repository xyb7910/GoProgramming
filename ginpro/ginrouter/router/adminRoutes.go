package router

import (
	"ginpro/ginrouter/middleware"
	"github.com/gin-gonic/gin"
)

// AdminRoutersInit  管理后台路由
func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middleware.InitMiddleWareOne, middleware.InitMiddleWareTwo)
	{
		adminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "后台首页")
		})
		adminRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "用户列表")
		})
		adminRouters.GET("/article", func(c *gin.Context) {
			c.String(200, "新闻列表")
		})
	}
}
