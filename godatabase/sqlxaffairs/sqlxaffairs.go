package main

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// define global variable
var db *sqlx.DB

// TransactionDemo demo transaction
func TransactionDemo() (err error) {
	tx, err := db.Beginx()
	if err != nil {
		fmt.Printf("begin transaction error: %v\n", err)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			fmt.Println("rollback transaction")
			tx.Rollback()
		} else {
			err = tx.Commit()
			fmt.Println("commit success")
		}
	}()

	sqlStr1 := "UPDATE user SET age = age + 1 WHERE id = ?"

	ret1, err := tx.Exec(sqlStr1, 1)
	if err != nil {
		fmt.Printf("exec sql error: %v\n", err)
		return err
	}
	n1, err := ret1.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows error: %v\n", err)
		return err
	}
	fmt.Printf("sqlStr1 affected rows: %d\n", n1)
	if n1 != 1 {
		return errors.New("update age error")
	}

	sqlStr2 := "UPDATE user SET age = 50 WHERE id = ?"
	if err != nil {
		fmt.Printf("exec sql error: %v\n", err)
		return err
	}
	ret2, err := tx.Exec(sqlStr2, 2)
	if err != nil {
		fmt.Printf("exec sql error: %v\n", err)
		return err
	}
	n2, err := ret2.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows error: %v\n", err)
		return err
	}
	fmt.Printf("sqlStr2 affected rows: %d\n", n2)
	if n2 != 1 {
		return errors.New("update age error")
	}
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
	} else {
		fmt.Println("init db success")
	}

	TransactionDemo()
}
