package main

import (
	"fmt"
	"time"
)

//函数参数传递的时候，值传递
func add0(a, b int) {
	println(a + b)
}

func add(a, b int) int {
	return a + b
}
func add2(a, b int) (sum int, err error) { //可以直接定义返回值变量名
	sum = a + b
	return sum, err
}

func add1(desc string, items ...int) (sum int, err error) {
	for _, value := range items {
		sum += value
	}
	return sum, err
}

func cal(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("+++")
		}
	case "-":
		return func() {
			fmt.Println("---")
		}
	default:
		return func() {
			fmt.Println("what")
		}
	}
}

func cal1(myfunc func(items ...int) int) int {
	return myfunc()
}

func runForever() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("doing")
	}
}

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func autoIncrement() func() int {
	local := 0 //一个函数中访问另外一个函数的局部变量是不可行的，但是匿名函数可以实现
	return func() int {
		local += 1
		return local
	}
}

func main() {
	//go函数支持普通函数，匿名函数，闭包
	/*
		go 函数是 "一等公民"
		1.函数本身可以作为变量
		2.匿名函数 闭包
		3.函数可以满足接口
	*/
	sum, _ := add2(1, 2)
	sum1, _ := add1("hahahh", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sum1)
	fmt.Println(sum)

	//函数作为一个变量
	funVar := add
	res := funVar(1, 2)
	fmt.Println(res)

	cal("+", 1, 3, 3, 4)() //返回值为函数必须调用

	//函数的参数也是一个函数
	cal1(func(items ...int) int {
		sum := 0
		for _, value := range items {
			sum += value
		}
		return sum
	})

	callback(2, add0)

	//没有函数名是为匿名函数
	callback(5, func(a, b int) {
		fmt.Println(a + b)
	})

	//函数闭包 实现一个函数每次调用一次返回的结果值都是增加一次后的值
	nectFunc := autoIncrement()
	for i := 0; i < 5; i++ {
		fmt.Println(nectFunc())
	}

	nectFunc1 := autoIncrement()
	for i := 0; i < 3; i++ {
		fmt.Println(nectFunc1())
	}
}
