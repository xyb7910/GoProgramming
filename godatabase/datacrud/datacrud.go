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

// QueryRowDemo query row demo
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id =?"
	var u User
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("query row error: %v\n", err)
		return
	}
	fmt.Printf("id: %d, age: %d, name: %s\n", u.id, u.age, u.name)
}

// QueryMultiRowDemo query multi row demo
func QueryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id>0"
	row, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query row error: %v\n", err)
		return
	}
	defer row.Close()

	for row.Next() {
		var u User
		err := row.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan row error: %v\n", err)
			return
		}
		fmt.Printf("id: %d, age: %d, name: %s\n", u.id, u.age, u.name)
	}
}

// InsertRowDemo insert row demo
func InsertRowDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "zml", 18)
	if err != nil {
		fmt.Printf("insert row error: %v\n", err)
		return
	}

	theID, err := ret.LastInsertId() // 获取最后插入的ID
	if err != nil {
		fmt.Printf("get last insert id error: %v\n", err)
		return
	}
	fmt.Printf("last insert id: %d\n", theID)
}

// UpdateRowDemo update row demo
func UpdateRowDemo() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := db.Exec(sqlStr, 20, 1)
	if err != nil {
		fmt.Printf("update row error: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响行数
	if err != nil {
		fmt.Printf("get rows affected error: %v\n", err)
		return
	}
	fmt.Printf("rows affected: %d\n", n)
}

// DeleteRowDemo delete row demo
func DeleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 4)

	if err != nil {
		fmt.Printf("delete row error: %v\n", err)
		return
	}

	n, err := ret.RowsAffected() //操作影响行数
	if err != nil {
		fmt.Printf("get rows affected error: %v\n", err)
		return
	}
	fmt.Printf("rows affected: %d\n", n)
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

	// query row demo
	//QueryRowDemo()

	// query multi row demo
	//QueryMultiRowDemo()

	// insert row demo
	//InsertRowDemo()

	// update row demo
	//UpdateRowDemo()

	// delete row demo
	//DeleteRowDemo()
}
