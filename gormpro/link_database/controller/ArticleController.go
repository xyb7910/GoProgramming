package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gormpro/link_database/models"
)

type ArticleController struct {
}

func (a ArticleController) Add(c *gin.Context) {
	// 批量插入
	articles := []*models.Article{
		{
			Model:       gorm.Model{ID: 1},
			Title:       "go 语言介绍",
			Content:     "Go是谷歌开发的一款编程语言...",
			AuthorID:    1,
			IsPublished: true,
			CategoryID:  1,
		},
		{
			Model:       gorm.Model{ID: 2},
			Title:       "python 语言介绍",
			Content:     "Python是一种广泛使用的高级编程语言...",
			AuthorID:    1,
			IsPublished: true,
			CategoryID:  1,
		},
		{
			Model:       gorm.Model{ID: 3},
			Title:       "go 语言介绍",
			Content:     "Go是谷歌开发的一款编程语言...",
			AuthorID:    2,
			IsPublished: true,
			CategoryID:  1,
		},
		{
			Model:       gorm.Model{ID: 4},
			Title:       "python 语言介绍",
			Content:     "Python是一种广泛使用的高级编程语言...",
			AuthorID:    2,
			IsPublished: true,
			CategoryID:  1,
		},
		{
			Model:       gorm.Model{ID: 5},
			Title:       "java 语言介绍",
			Content:     "java是一种复杂的高级编程语言...",
			AuthorID:    2,
			IsPublished: true,
			CategoryID:  1,
		},
		{
			Model:       gorm.Model{ID: 6},
			Title:       "python 语言介绍",
			Content:     "Python是一种广泛使用的高级编程语言...",
			AuthorID:    3,
			IsPublished: true,
			CategoryID:  1,
		},
	}
	models.DB.Create(&articles)
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}

func (a ArticleController) Get(c *gin.Context) {
	article := &[]models.Article{}
	//  < 小于号， 选择所有id小于3的文章
	models.DB.Debug().Where("id < ?", 3).Find(&article)
	//  SELECT * FROM `articles` WHERE id < 3 AND `articles`.`deleted_at` IS NULL

	// > 大于号， 选择所有id大于5的文章
	var id = 5
	models.DB.Debug().Where("id > ?", id).Find(&article)
	// SELECT * FROM `articles` WHERE id > 5 AND `articles`.`deleted_at` IS NULL

	// AND 操作符， 选择所有id大于1且小于5的文章
	models.DB.Debug().Where("id > ? and id < ?", 1, 5).Find(&article)
	//  SELECT * FROM `articles` WHERE (id > 1 AND id < 5) AND `articles`.`deleted_at` IS NULL

	// in 操作符， 选择所有id大于2且小于5的文章
	models.DB.Debug().Where("id in ?", []int{3, 4}).Find(&article)
	//SELECT * FROM `articles` WHERE id in (3,4) AND `articles`.`deleted_at` IS NULL

	// like 操作符， 选择 文章标题含 go 的文章
	models.DB.Debug().Where("title like ?", "%go%").Find(&article)
	// SELECT * FROM `articles` WHERE title like '%go%' AND `articles`.`deleted_at` IS NULL

	// between and 操作符， 选择文章id在 1 到 5 之间的文章
	models.DB.Debug().Where("id between ? and ?", 1, 5).Find(&article)
	// SELECT * FROM `articles` WHERE (id between 1 and 5) AND `articles`.`deleted_at` IS NULL

	// or 操作符， 选择文章id在 1 到 5 之间的文章或 id 等于 6 的文章
	models.DB.Debug().Where("id between ? and ? or id = ?", 1, 5, 6).Find(&article)
	//SELECT * FROM `articles` WHERE (id between 1 and 5 or id = 6) AND `articles`.`deleted_at` IS NULL
	models.DB.Debug().Where("id between ? and ?", 1, 5).Or("id = ?", 6).Find(&article)
	//SELECT * FROM `articles` WHERE ((id between 1 and 5) OR id = 6) AND `articles`.`deleted_at` IS NULL

	c.JSON(200, gin.H{
		"success": true,
		"result":  article,
	})
}

func (a ArticleController) GetOne(c *gin.Context) {
	article := &[]models.Article{}

	// 查询文章的id 和 title
	models.DB.Debug().Select("id, title").Find(&article)
	// SELECT id, title FROM `articles` WHERE `articles`.`deleted_at` IS NULL

	// 按 ID 降序排序
	models.DB.Debug().Order("id desc").Find(&article)
	// SELECT * FROM `articles` WHERE `articles`.`deleted_at` IS NULL ORDER BY id desc

	// 按 ID 升序排序 按 author_id 降序排序
	models.DB.Debug().Order("id asc, author_id desc").Find(&article)
	// SELECT * FROM `articles` WHERE `articles`.`deleted_at` IS NULL ORDER BY id asc, author_id desc

	// 查询ID大于1的文章， 限制1条数据 （随机一条）
	models.DB.Debug().Where("id > ?", 1).Limit(1).Find(&article)
	// SELECT * FROM `articles` WHERE id > 1 AND `articles`.`deleted_at` IS NULL LIMIT 1

	// 实现 跳过一条数据查询2条数据 (分页)
	models.DB.Debug().Where("id > ?", 1).Offset(1).Limit(1).Find(&article)
	// SELECT * FROM `articles` WHERE id > 1 AND `articles`.`deleted_at` IS NULL LIMIT 1 OFFSET 1
	c.JSON(200, gin.H{
		"success": true,
		"result":  article,
	})
}

func (a ArticleController) GetSum(c *gin.Context) {
	article := &[]models.Article{}

	var sum int64
	// 获取author_id等于1的文章的数量
	models.DB.Debug().Where("author_id = ?", 1).Find(&article).Count(&sum)
	//SELECT count(*) FROM `articles` WHERE author_id = 1 AND `articles`.`deleted_at` IS NULL

	c.JSON(200, gin.H{
		"success": true,
		"result":  article,
	})
}

func (a ArticleController) GetDistinct(c *gin.Context) {
	article := &[]models.Article{}

	// 选择不同的 title
	models.DB.Debug().Distinct("title").Find(&article)
	// SELECT DISTINCT `title` FROM `articles` WHERE `articles`.`deleted_at` IS NULL
	c.JSON(200, gin.H{
		"success": true,
		"result":  article,
	})
}

type Result struct {
	AuthorId int
	Title    string
}

func (a ArticleController) GetScan(c *gin.Context) {
	var result []Result
	//models.DB.Table("articles").Select("author_id, title").Scan(&result)
	models.DB.Raw("SELECT author_id, title FROM articles").Scan(&result)
	fmt.Println(result)
}
