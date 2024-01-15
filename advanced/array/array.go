package main

import "fmt"

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{2, 3, 4, 5, 6}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	printArray(arr1)
	//printArray(arr2) // go语言中数组为一个数，类型不同时不能使用
	printArray(arr3)
}
