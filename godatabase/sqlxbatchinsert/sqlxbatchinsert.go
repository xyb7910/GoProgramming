package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// define global variable
var db *sqlx.DB

// BatchInsertUsers 自己构造一个批量插入的函数
func BatchInsertUsers(users []*User) (err error) {
	// 存放 (?, ?) 的 slice
	valueStrings := make([]string, 0, len(users))
	// 存放 values 的 slice
	valueArgs := make([]interface{}, 0, len(users))
	// 遍历 users
	for _, user := range users {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, user.Name)
		valueArgs = append(valueArgs, user.Age)
	}
	// 自行拼接成 sql
	stmt := fmt.Sprintf("INSERT INTO users (name, age) VALUES %s",
		strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)
	return err
}

// Value 实现 sql/driver.Valuer 接口
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers2 使用sqlx自带的批量插入函数
func BatchInsertUsers2(users []interface{}) (err error) {
	query, args, _ := sqlx.In(
		"insert into user (name, age) values (?), (?), (?)",
		users...,
	)
	fmt.Println(query)
	fmt.Println(args)
	_, err = db.Exec(query, args...)
	return err
}

// InitDB initialize database connection
func InitDB() (err error) {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect mysql error: %v\n", err)
		return
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	return
}

func main() {
	err := InitDB()
	if err != nil {
		fmt.Printf("init db error: %v\n", err)
		return
	} else {
		fmt.Println("init db success")
	}

}
