package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

// 定义一些正则表达式常量
const (
	EmailRegexPattern    = `^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$`
	PasswordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// UserHandler 定义 user 相关的 handler
type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

// NewUserHandler 实现 user 相关的 handler 的构造函数
func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailRegexp:    regexp.MustCompile(EmailRegexPattern),
		passwordRegexp: regexp.MustCompile(PasswordRegexPattern),
	}
}

// RegisterRoute 实现路由注册
func (u *UserHandler) RegisterRoute(server *gin.Engine) {
	// 首先进行路由分组
	user := server.Group("/user")

	// 定义其他路由
	user.POST("/signup", u.SignUp)
	user.POST("/signin", u.SignIn)
	user.GET("/profile", u.Profile)
	user.GET("/logout", u.Logout)
	user.GET("/edit", u.Edit)
}

// SignUp 实现 user 相关的 signup 接口
func (u *UserHandler) SignUp(context *gin.Context) {
	type SignUpRequest struct {
		Username        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}
	var request SignUpRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		return
	}

	// 邮箱格式校验
	isEmail := u.emailRegexp.MatchString(request.Email)
	if !isEmail {
		context.JSON(http.StatusBadRequest, "your input is not email")
		return
	}

	// 密码格式校验
	isPassword := u.passwordRegexp.MatchString(request.Password)
	if !isPassword {
		context.JSON(http.StatusBadRequest, "your input is not password")
	}

	// 检验两次密码是否一致
	if request.Password != request.ConfirmPassword {
		context.JSON(http.StatusBadRequest, "password not match")
		return
	}
}

// SignIn 实现 user 相关的 signin 接口
func (u *UserHandler) SignIn(context *gin.Context) {

}

// Profile 实现 user 相关的 profile 接口
func (u *UserHandler) Profile(context *gin.Context) {

}

// Logout 实现 user 相关的 logout 接口
func (u *UserHandler) Logout(context *gin.Context) {

}

// Edit 实现 user 相关的 edit 接口
func (u *UserHandler) Edit(context *gin.Context) {

}
