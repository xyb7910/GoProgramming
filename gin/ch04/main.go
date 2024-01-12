package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   int    `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()

	router.GET("/:name/:id", handler)
	router.Run(":8080")
}

func handler(context *gin.Context) {
	var person Person
	if err := context.ShouldBindUri(&person); err != nil { //约束url
		context.JSON(404, gin.H{"msg": err})
		return
	}
	context.JSON(200, gin.H{
		"name": person.Name,
		"uuid": person.ID,
	})
}
