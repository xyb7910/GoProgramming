package main

import "fmt"

func main() {
	var c1 byte //主要用来存放字符类型
	var c2 rune //也是字符
	c1 = 'a'
	c2 = 'c'
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Printf("a=%c\n", c1)
	fmt.Printf("c=%c\n", c2)

	var name string
	name = "YPB"
	fmt.Println(name)
}
