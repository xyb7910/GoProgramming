package main

import "fmt"

func main() {
	// 运算符
	var a, b = 7, 2
	var str1, str2, str3 = "hello", " ", "world"
	fmt.Println(a + b)
	fmt.Println(str1 + str2 + str3)
	fmt.Println(a % b)
	fmt.Println(a / b)
	a++
	b--
	fmt.Println(a + b)
	fmt.Println((a + b) >> 1)
}
