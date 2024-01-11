package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type Language struct {
	gorm.Model
	Name    string
	AddTime time.Time //每个记录创建的时候自动加上当前时间加入到AddTime中
}

func (l *Language) BrforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return
}

// 通过给某一个struct添加TableName方法来自定义表名

//func (Language) TableName() string {
//	return "my_language"
//}

/*
1、我们自己定义表名是什么
2、统一的给所有的表名加上一个前缀
*/

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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//NamingStrategy 和 TableName 不能同时配置
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "mxshop_",
		},
		Logger: newLogger})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "go",
	})
}
