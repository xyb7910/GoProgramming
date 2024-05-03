package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

type Message struct {
	OutputLang string `json:"output_lang"`
	Text       string `json:"text"`
}

func HandleMessage(c *gin.Context) {
	var Message Message
	if err := c.BindJSON(&Message); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	prompt := prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("你是一个翻译机器人", nil),
		prompts.NewHumanMessagePromptTemplate(`翻译这段文字到{{.OutputLang}}: {{.Text}}`, []string{".OutputLang", ".Text"}),
	})

	// 填充数据
	vals := map[string]any{
		"OutputLang": Message.OutputLang,
		"Text":       Message.Text,
	}

	messages, err := prompt.FormatMessages(vals)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 链接ollama
	llm, err := ollama.New(ollama.WithModel("qwen"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	content := []llms.MessageContent{
		llms.TextParts(messages[0].GetType(), messages[0].GetContent()),
		llms.TextParts(messages[1].GetType(), messages[1].GetContent()),
	}

	generateContent, err := llm.GenerateContent(context.Background(), content)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"result": generateContent.Choices[0].Content})
}

func main() {
	r := gin.Default()

	r.POST("/ollama", HandleMessage, func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	r.Run(":8080")
}
