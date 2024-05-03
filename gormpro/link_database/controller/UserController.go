package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gormpro/link_database/models"
	"time"
)

type UserController struct {
}

func (con UserController) Add(c *gin.Context) {
	// 单数据
	user := models.User{
		UserName: "ypc",
		Age:      20,
		Email:    "ychag@example.com",
		AddTime:  int(time.Now().Unix()),
	}

	// 批量数据
	//users := []*models.User{
	//	{UserName: "yxc", Age: 20, Email: "ychag@example.com", AddTime: int(time.Now().Unix())},
	//	{UserName: "zmz", Age: 20, Email: "zmzag@example.com", AddTime: int(time.Now().Unix())},
	//}

	// INSERT INTO `user` (`user_name`,`age`,`email`,`add_time`) VALUES ('ypc',20,'ychag@example.com',1714534305)
	result := models.DB.Debug().Create(&user) // 通过指针创建

	if result.RowsAffected > 1 {
		fmt.Println(user.Id)
	}

	fmt.Println(result.RowsAffected)
	//fmt.Println(user.Id)
	c.String(200, "add ok")
}

func (con UserController) Get(c *gin.Context) {
	user := models.User{}

	// First 获取第一条记录（主键升序）
	models.DB.Debug().First(&user)
	// SELECT * FROM `user` ORDER BY `user`.`id` LIMIT 1
	fmt.Println(user.Id)

	// Last 获取最后一条记录（主键降序）
	models.DB.Debug().Last(&user)
	fmt.Println(user.Id)
	// SELECT * FROM `user` WHERE `user`.`id` = 1 ORDER BY `user`.`id` DESC LIMIT 1

	// Find 获取所有记录
	users := []models.User{}
	models.DB.Debug().Find(&users)
	//SELECT * FROM `user`

	c.JSON(200, gin.H{
		"message": "get ok",
		"result":  users,
	})
}

func (con UserController) GetByCondition(c *gin.Context) {
	user := models.User{}

	// Where 指定查询条件
	models.DB.Debug().Where("user_name = ?", "yxc").First(&user)
	//  SELECT * FROM `user` WHERE user_name = 'yxc' ORDER BY `user`.`id` LIMIT 1
	c.JSON(200, gin.H{
		"message": "get ok",
		"result":  user,
	})
}

func (con UserController) Update(c *gin.Context) {

	// 首先获取数据
	user := models.User{Id: 1}
	models.DB.Debug().Find(&user)
	//  SELECT * FROM `user` WHERE `user`.`id` = 1

	// 然后更新数据
	user.UserName = "haha"
	user.Age = 22
	models.DB.Debug().Save(&user)
	// UPDATE `user` SET `user_name`='haha',`age`=22,`email`='ychag@example.com',`add_time`=1708047687 WHERE `id` = 1

	c.String(200, "update ok")
}

func (con UserController) UpdateByCondition(c *gin.Context) {
	// Where 条件更新单列数据
	models.DB.Debug().Model(&models.User{}).Where("id = ?", 1).Update("age", 15)
	// UPDATE `user` SET `age`=15 WHERE id = 1

	// Where 条件更新多列数据
	var users []models.User
	models.DB.Debug().Model(&users).Where("user_name = ?", "yxc").Updates(models.User{UserName: "yxc", Age: 22})
	// UPDATE `user` SET `user_name`='yxc',`age`=22 WHERE user_name = 'yxc'
	c.String(200, "update ok")
}

func (con UserController) Delete(c *gin.Context) {
	// Where 条件删除数据
	models.DB.Debug().Where("user_name = ?", "haha").Delete(&models.User{})
	// DELETE FROM `user` WHERE user_name = 'haha'
}

func (con UserController) DeleteAll(c *gin.Context) {
	users := []models.User{}
	models.DB.Debug().Where("id > ?", 0).Delete(&users)
	//  SELECT * FROM `user` WHERE id > 0
	c.JSON(200, gin.H{
		"message": "delete ok",
	})
}
