package main

import (
	"LearingGo/ginpro/ginsession/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 创建一个基于 cookie 的 session store
	// secret  密钥
	//store := cookie.NewStore([]byte("secret111111"))

	//初始化基于redis的存储引擎: 需要启动redis服务,不然会报错
	//参数说明:
	//自第1个参数-redis最大的空闲连接数
	//第2个参数-数通信协议tcp或者udp
	//第3个参数-redis地址,格式，host:port 第4个参数-redis密码
	//第5个参数-session加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	//设置session中间件，参数my_session，指的是session的名字，也是cookie的名字
	//store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("my_session", store))

	r.GET("/set", controller.DefaultController{}.SetSession, func(c *gin.Context) {
		c.String(200, "set ok")
	})
	r.GET("/get", controller.DefaultController{}.GetSession, func(c *gin.Context) {
		c.String(200, "get ok")
	})

	r.Run(":8080")
}
