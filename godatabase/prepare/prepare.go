package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id   int
	name string
	age  int
}

// 定义全局数据库
var db *sql.DB

// PrepareQueryDemo prepare query demo
func PrepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error: %v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(0)

	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan error: %v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}

}

// PrepareInsertDemo prepare insert demo
func PrepareInsertDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error: %v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec("jack", 18)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}

	_, err = stmt.Exec("Amy", 18)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}

	fmt.Println("insert success")
}

// PrepareUpdateDemo prepare update demo
func PrepareUpdateDemo() {
	sqlStr := "update user set age = ? where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error: %v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(20, 6)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}

	_, err = stmt.Exec(20, 7)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}
	fmt.Println("update success")
}

// PrepareDeleteDemo prepare delete demo
func PrepareDeleteDemo() {
	sqlStr := "delete from user where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare error: %v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(8)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}

	_, err = stmt.Exec(7)
	if err != nil {
		fmt.Printf("exec error: %v\n", err)
		return
	}
	fmt.Println("delete success")
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
	err := InitDB()

	if err != nil {
		fmt.Printf("init db error: %v\n", err)
		return
	} else {
		fmt.Println("link db success")
	}

	defer db.Close()

	// PrepareQueryDemo()
	PrepareQueryDemo()

	// PrepareInsertDemo()
	//PrepareInsertDemo()

	// PrepareUpdateDemo()
	//PrepareUpdateDemo()

	// PrepareDeleteDemo()
	PrepareDeleteDemo()

	PrepareQueryDemo()

}
