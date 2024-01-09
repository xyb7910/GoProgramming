package main

import (
	"fmt"
)

func main() {
	//go语言中的集合类型的数据结构，数组，切片，map，list

	/*
		//数组： 定义 var name[count] int
		var courses1 [3]string // 只有三个元素的数组类型，Go语言的定义很严格
		//var courses2 [4]string
		//注意courses1 与courses2 为不同类型的数组

		//fmt.Printf("%T%\r\n", courses1)
		//fmt.Printf("%T", courses2)

		courses1[0] = "go"
		courses1[1] = "grpc"
		courses1[2] = "gin"

		fmt.Println(courses1)

		for _, value := range courses1 {
			fmt.Println(value)
		}
	*/

	/*
		//数组的初始化
		courses2 := [3]string{"go", "grpc", "gin"} //最常用
		for _, value := range courses2 {
			fmt.Println(value)
		}
		fmt.Println("*********")
		for i := 0; i < len(courses2); i++ {
			fmt.Println(courses2[i])
		}
		courses3 := [...]string{2: "gin"}
		for _, value := range courses3 {
			fmt.Println(value)
		}
		courses4 := [...]string{"go", "gin"} //...表示长度为初始化
		for _, value := range courses4 {
			fmt.Println(value)
		}
	*/

	//多维数组
	var courseInfo [3][4]string
	courseInfo[0] = [4]string{"go", "1h", "yxc", "Alice"}
	courseInfo[1] = [4]string{"c++", "2h", "yxc", "Tom"}
	courseInfo[2] = [4]string{"go", "1h", "yxc", "google"}

	fmt.Println(len(courseInfo))

	for i := 0; i < len(courseInfo); i++ {
		for j := 0; j < len(courseInfo[i]); j++ {
			fmt.Print(courseInfo[i][j] + " ")
		}
		fmt.Println()
	}

	for _, row := range courseInfo {
		fmt.Println(row)
	}
}
