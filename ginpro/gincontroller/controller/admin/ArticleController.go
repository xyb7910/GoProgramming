package admin

import "github.com/gin-gonic/gin"

type ArticleController struct{}

func (a ArticleController) GetArticle(c *gin.Context) {
	c.JSON(200, "获取文章")
}

func (a *ArticleController) AddArticle(c *gin.Context) {
	c.JSON(200, "添加文章")
}

func (a *ArticleController) EditArticle(c *gin.Context) {
	c.JSON(200, "编辑文章")
}
