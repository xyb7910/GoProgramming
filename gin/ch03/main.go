package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化一个gin的server对象
	router := gin.Default()

	/*
		v1 := router.Group("/v1")
		{
			v1.GET("/login", loginEndpoint1)
			v1.GET("/register", registerEndpoint1)
			v1.GET("/read", readEndpoint1)
		}
	*/
	//Simple group 1

	//Simple group 2
	v2 := router.Group("/v2") //路由分组
	{
		v2.GET("/", loginEndpoint2)
		v2.GET("/:register/:action", registerEndpoint2) //带参数的url
		v2.GET("/read/*who", readEndpoint2)             //获取路由分组的参数
	}

	router.Run(":8080")
}

func readEndpoint2(context *gin.Context) {
	context.JSON(200, gin.H{
		"action": "post",
	})
}

func registerEndpoint2(context *gin.Context) {
	register := context.Param("register")
	action := context.Param("action")
	context.JSON(200, gin.H{
		"register": register,
		"action":   action,
	})
}

func loginEndpoint2(context *gin.Context) {
	context.JSON(200, gin.H{
		"name": "hahaha",
	})
}

/*
func registerEndpoint1(context *gin.Context) {

}

func readEndpoint1(context *gin.Context) {

}

func loginEndpoint1(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "ypb",
	})
}
*/
