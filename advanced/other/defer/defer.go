package main

import "fmt"

func dosomething() {
	defer fmt.Println("do somethings test2")
	panic("panic1")

	fmt.Println("do something2")

}

func main() {

	fmt.Println("go somethings1")
	dosomething()
	defer fmt.Println("do somethings test1")
}
