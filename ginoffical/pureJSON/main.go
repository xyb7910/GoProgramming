package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// use unicode
	r.GET("/unicode", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>hello</b>",
		})
	})

	// use pure json
	r.GET("/pure", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>hello</b>",
		})
	})
	r.Run(":8080")
}
