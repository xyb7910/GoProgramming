package main

import "fmt"

func main() {
	/*
		//定义常量，定义时指定值，不能修改， 常量定义时尽量全部大写
		const PI float32 = 3.1415926 //显示定义
		const PI = 3.1415926 //隐式定义
		fmt.Println(PI）
	*/

	const (
		AGE  = 18
		NAME = "YPB"
		ADDR = "SHANXI"
	)

	const (
		Z int = 18
		Y
		C = "S"
		D
	)
	fmt.Println(Z, Y, C, D)
}
