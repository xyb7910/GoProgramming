package main

import "fmt"

func updateSlice(s []int) { //slice可以改变值
	s[0] = 100
}

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	s1 := arr[2:6]
	fmt.Println(s1)
	updateSlice(s1)
	fmt.Println(s1)
}
