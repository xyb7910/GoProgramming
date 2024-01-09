package main

import (
	"fmt"
	"strconv"
)

func printSlice(data []string) {
	data[0] = "java"

	for i := 0; i < 10; i++ {
		data = append(data, strconv.Itoa(i))
	}
	fmt.Println(data)
}

func main() {
	//go的slice在函数传递的时候是 值传递，效果上又呈现了引用传递的效果
	/*
		course := []string{"go", "grpc", "gin"}
		printSlice(course)
		fmt.Println(course)
	*/

	/*
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s1 := data[1:6]
		s2 := data[2:7]

		fmt.Println(len(s1), cap(s1))
		fmt.Println(len(s2), cap(s2))

		s2 = append(s2, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2)

		s2[0] = 22

		fmt.Println(s1) // 2 3 4 5 6
		fmt.Println(s2) //22 4 5 6 7 1 1 1 1 1 1 1 1 2 2 2
		fmt.Println(len(s1), cap(s1))
		fmt.Println(len(s2), cap(s2))
	*/

	var data []int
	for i := 0; i < 2000; i++ {
		data = append(data, i)
		fmt.Println(len(data), cap(data))
	}
}
