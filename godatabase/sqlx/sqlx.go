package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// define global var
var db *sqlx.DB

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// QueryRowDemo query row demo
func QueryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u User
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// QueryMultiRowDemo query multi row demo
func QueryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []User
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("users:%v\n", users)
}

// InsertRoweDemo insert demo
func InsertRoweDemo() {
	sqlStr := "insert into user(name,age) values(?,?)"
	ret, err := db.Exec(sqlStr, "toby", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // get last insert id
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("last insert id:%d\n", theID)
}

// UpdateRowDemo update demo
func UpdateRowDemo() {
	sqlStr := "update user set age=? where id =?"
	ret, err := db.Exec(sqlStr, 20, 1)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // get affected rows
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
		return
	}
	fmt.Printf("affected rows:%d\n", n)
}

func DeleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // get affected rows
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
		return
	}
	fmt.Printf("affected rows:%d\n", n)
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

	// QueryRowDemo()
	//InsertRoweDemo()
	//UpdateRowDemo()
	//DeleteRowDemo()
	QueryMultiRowDemo()
}
