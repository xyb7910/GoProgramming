package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	name    string
	age     int32
	address string
}

func main() {
	/*
		p1 := Person{}
		//Scan 家族 : 以空格为分隔符分割标准输入中的内容，并将分割后的内容保存在给定的变量中

		fmt.Scanf("%s%d%s", &p1.name, &p1.age, &p1.address)
		fmt.Printf("你的名字是：%s\n", p1.name)
		fmt.Printf("你的年龄是：%d\n", p1.age)
		fmt.Printf("你的地址是：%s\n", p1.address)

	*/

	/*
		//SScan家族 从给定字符串中读取数据
		var (
			name string
			age  int32
		)
		input := "ypb 18"
		fmt.Sscan(input, &name, &age)
		fmt.Printf("name: %s\t, age:%d\n", name, age)
	*/

	//使用bufio
	var inputReader *bufio.Reader
	var input string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入姓名: ")
	input, err = inputReader.ReadString(' ')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}

}
