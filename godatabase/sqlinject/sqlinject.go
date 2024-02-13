package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// define global variable
var db *sql.DB

type User struct {
	id   int
	age  int
	name string
}

// SqlInject sql inject
func SqlInject(sql string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", sql)
	fmt.Printf("sql: %s\n", sqlStr)

	var user User
	err := db.QueryRow(sqlStr).Scan(&user.id, &user.name, &user.age)
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}
	fmt.Printf("user: %v\n", user)
}

// InitDB initialize database
func InitDB() (err error) {
	dsn := "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 初始化数据库
	// Open 打开数据库连接,返回一个 *sql.DB 实例,并不会去连接数据库
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// 连接数据库
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// call InitDB
	err := InitDB()

	if err != nil {
		fmt.Printf("init db error: %v\n", err)
		return
	} else {
		fmt.Println("link db success")
	}

	defer db.Close()

	// call SqlInject
	// 以下都可以触发sql注入
	SqlInject("xxx' or 1=1#")
	SqlInject("xxx' union select * from user #")
	SqlInject("xxx' and (select count(*) from user) <10 #")

}
