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

// User 拥有并属于多种 language，`user_languages` 是连接表
type User3 struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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

	//db.AutoMigrate(&User3{})
	//
	//languages := []Language{}
	//languages = append(languages, Language{Name: "go"})
	//languages = append(languages, Language{Name: "java"})
	//
	//user := User3{
	//	Languages: languages,
	//}
	//
	//db.Create(&user)

	var user User3
	//db.Preload("Languages").First(&user)
	//for _, language := range user.Languages {
	//	fmt.Println(language.Name)
	//}

	db.First(&user)
	var languages []Language
	_ = db.Model(&user).Association("Languages").Find(&languages)
	for _, language := range languages {
		fmt.Println(language.Name)
	}
}
