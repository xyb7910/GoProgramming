package models

import (
	"gorm.io/gorm"
)

type User struct {
	// 默认表名是 users
	Id       int
	UserName string
	Age      int
	Email    string
	AddTime  int
}

type Article struct {
	gorm.Model
	Title       string
	Content     string
	AuthorID    uint
	IsPublished bool
	CategoryID  uint
}

func (User) TableName() string {
	return "user"
}
