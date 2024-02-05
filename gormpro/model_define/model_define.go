package model_define

import (
	"gorm.io/gorm"
	"time"
)

/*
 gorm 定义偏向与约定，而不是配置
 默认情况：
	使用 ID 作为主键
	使用结构体名的 蛇形 作为表名
	字段名使用 蛇形 作为列名
	使用 CreateAt, UpdateAt 追踪创建和更新时间
*/

// gorm.Model

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatAt   time.Time
	UptAt     time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 嵌入结构体

type User struct {
	gorm.Model
	Name string
}

// 等效于

//type User struct {
//	ID        uint `gorm:"primary_key"`
//	CreatAt   time.Time
//	UptAt     time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	Name string
//}

// 使用 标签 embeded 来指定 嵌入结构体
type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID     int    `gorm:"primary_key"`
	Author Author `gorm:"embedded"`
}

// 等效于

//type Blog struct {
//	ID int `gorm:"primary_key"`
//	Name string
//	Email string
//}

// 可以使用 embeddedPrefix 来指定 嵌入结构体的前缀

//type Blog struct {
//	ID int `gorm:"primary_key"`
//	Author Author `gorm:"embedded;embeddedPrefix:author_"`
//}
