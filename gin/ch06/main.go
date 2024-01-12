package main

import "github.com/gin-gonic/gin"
import "LearingGo/gin/ch06/proto"

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProto)
	router.Run(":8080")
}

func returnProto(context *gin.Context) {
	course := []string{"python", "go", "c++"}
	user := &proto.Teacher{
		Name:   "ypb",
		Course: course,
	}
	context.ProtoBuf(200, user)
}

func moreJSON(context *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "ypb"
	msg.Message = "hello"
	msg.Number = 123

	context.JSON(200, msg)
}
