package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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
	db.AutoMigrate()
}
