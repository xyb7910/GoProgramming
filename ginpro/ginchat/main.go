package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms/ollama"
)

func HandleCaht(c *gin.Context) {
	var RequestDate struct {
		Prompt string `json:"prompt"`
	}

	if err := c.BindJSON(&RequestDate); err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("invalid request: %v", err)})
		return
	}

	llm, err := ollama.New(ollama.WithModel("qwen"))
	if err != nil {
		c.JSON(400, gin.H{"error": "create llm error"})
	}

	call, err := llm.Call(context.Background(), RequestDate.Prompt)
	if err != nil {
		c.JSON(400, gin.H{"error": "call error"})
	}
	c.JSON(200, gin.H{"response": call})
}

func main() {
	r := gin.Default()
	r.POST("/chat", HandleCaht)
	r.Run(":8080")
}
