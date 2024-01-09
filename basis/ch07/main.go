package main

import "fmt"

func main() {
	//type 关键字

	/*
		1. 定义结构体
		2. 定义接口
		3. 定义类型别名
		4. 定义类型
	*/
	type MyInt int
	var a MyInt = 12
	b := 12
	c := 13
	fmt.Println(b + c)
	fmt.Printf("%T\t%d\n", a, a) //类型为自定义类型 不能与其他类型进行运算
	fmt.Printf("%T\t%d\n", b, b) //类型为int
}
