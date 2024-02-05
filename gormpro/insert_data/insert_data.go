package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	Name     string
	Age      int
	Birthday time.Time
}

// mysql link database

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// Globally mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 迁移模型，新建表
	db.AutoMigrate(&User{})

	// 单一数据插入
	//user := User{
	//	Name:     "Leo",
	//	Age:      18,
	//	Birthday: time.Now(),
	//}
	//user1 := User{
	//	Name:     "Lei",
	//	Age:      18,
	//	Birthday: time.Now(),
	//}
	//user2 := User{
	//	Name:     "Yan",
	//	Age:      18,
	//	Birthday: time.Now(),
	//}
	//
	//// 通过数据指针，插入数据
	//db.Create(&user)
	//
	////创建记录并更新给出的字段
	//db.Select("Name", "Age", "CreatedAt").Create(&user1)
	//// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	//
	//// 创建记录并更新未给出的字段
	//db.Omit("Name", "Age", "CreatedAt").Create(&user2)
	//// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")

}
