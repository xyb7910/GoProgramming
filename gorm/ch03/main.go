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
	Name         string
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

	// 迁移 schema
	_ = db.AutoMigrate(&User{})

	user := User{Name: "yxc"}

	fmt.Println(user.ID)

	res := db.Create(&user)

	fmt.Println(user.ID)          // 返回插入数据的主键
	fmt.Println(res.Error)        //返回error
	fmt.Println(res.RowsAffected) //返回插入记录的条数

	//db.Model(&User{ID: 1}).Update("Name", "")
	//updates语句不会更新零值，但是update语句会更新
	db.Model(&User{ID: 1}).Updates(User{Name: ""})

	//解决仅更新非零值的两种方式
	/*
		1、将string 设置为 *string
		2、使用sql 中的 NULL*** 来解决
	*/
}
