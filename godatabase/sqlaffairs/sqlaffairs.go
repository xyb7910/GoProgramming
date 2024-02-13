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

// 确保两次更新操作要么同时成功要么同时失败，不会存在中间状态。

// TransactionDemo handel transaction
func TransactionDemo() {
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("transactionDemo error: %v\n", err)
		return
	}

	sqlStr1 := "UPDATE user SET age = age + 1 WHERE id = ?"

	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback()
		fmt.Printf("e1 error: %v\n", err)
		return
	}

	affectRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("affecRow1 error: %v\n", err)
		return
	}
	fmt.Printf("affectRow1: %d\n", affectRow1)

	sqlStr2 := "UPDATE user SET age = 40 WHERE id = ?"
	ret2, err := tx.Exec(sqlStr2, 2)
	if err != nil {
		fmt.Printf("ret2 error: %v\n", err)
		return
	}
	affectRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("affecRow2 error: %v\n", err)
		return
	}
	fmt.Printf("affecRow2: %d\n", affectRow2)

	if affectRow1 != 1 || affectRow2 != 1 {
		tx.Rollback()
		fmt.Println("transactionDemo failed")
	} else {
		tx.Commit()
		fmt.Println("transactionDemo success")
	}
	fmt.Println("transactionDemo end")
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

	TransactionDemo()
}
