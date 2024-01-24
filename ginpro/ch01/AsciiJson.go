package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	Field        string `form:"field_b"`
}

type StructC struct {
	NestedStructPoint *StructA
	Field             string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func main() {
	r := gin.Default()

	// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
	r.GET("/someJSON", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "go",
			"tag":  "<br>",
		}
		// {"lang":"go","tag":"\u003cbr\u003e"}
		context.AsciiJSON(http.StatusOK, data)
	})

	r.GET("/getb", GetDataB)

	r.Run(":8080")
}

func GetDataB(context *gin.Context) {
	var b StructB
	context.Bind(&b)
	context.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.Field,
	})
}
