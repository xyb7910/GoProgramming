package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// User 有多张 CreditCard，UserID 是外键

type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
	// 在大型高并发不建议使用外键约束影响效率， 但保持了数据的完整性，自己在业务层面保持数据的一致性
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
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

	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&User{})

	//user := User{}
	//db.Create(&user)
	//db.Create(&CreditCard{
	//	Number:    "12",
	//	UserRefer: user.ID,
	//})
	//db.Create(&CreditCard{
	//	Number:    "34",
	//	UserRefer: user.ID,
	//})

	var user User
	db.Preload("CreditCard").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card)
	}
}
