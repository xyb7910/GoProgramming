package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint
	MyName       string `gorm:"column:name"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(172.16.92.129:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}

	/*
			//通过first查询单个数据，获取第一条记录（主键升序）
			var user User
			result := db.First(&user, 2)
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				fmt.Println("未找到")
			}
			fmt.Println(user.ID)

		//通过主键查询
		var users []User
		result := db.First(&users, 22)
		fmt.Println("总记录数为：", result.RowsAffected)
		for _, user := range users {
			fmt.Println(user.ID)
		}
	*/

	//查询方式条件有三种：1.string 2.struct 3.map
	// string 条件
	//var user User
	var users []User
	// db.Where("name = ?", "yxc").First(&user)
	//db.Where(&User{MyName: "yxc", Age: 18}).First(&users) // 不必记住数据库表项的名字
	//SELECT * FROM `users` WHERE `users`.`name` = ? AND `users`.`age` = ? ORDER BY `users`.`id` LIMIT 1

	db.Where(map[string]interface{}{"name": "yxc", "Age": 18}).Find(&users) //必须与数据库同名
	//SELECT * FROM `users` WHERE `Age` = ? AND `name` = ?
	for _, user := range users {
		fmt.Println(user.ID)
	}
}
