package admin

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (con UserController) GetUserList(c *gin.Context) {
	//c.String(200, "用户列表1")
	con.success(c)
}

func (con UserController) AddUser(c *gin.Context) {
	c.String(200, "用户添加2")
}

func (con UserController) EditUser(c *gin.Context) {
	c.String(200, "用户编辑3")
}
