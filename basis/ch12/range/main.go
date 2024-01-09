package main

import (
	"fmt"
	"time"
)

func main() {
	var msg chan int

	msg = make(chan int, 4)

	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("all done")
	}(msg)
	msg <- 1 // 放值到channel中
	msg <- 2 // 放值到channel中
	msg <- 3 // 放值到channel中
	msg <- 4 // 放值到channel中

	close(msg)
	d1 := <-msg //虽然channel已经关闭，但仍可以取值，不能在放值了
	fmt.Println(d1)
	d2 := <-msg //虽然channel已经关闭，但仍可以取值，不能在放值了
	fmt.Println(d2)
	d3 := <-msg //虽然channel已经关闭，但仍可以取值，不能在放值了
	fmt.Println(d3)
	d4 := <-msg //虽然channel已经关闭，但仍可以取值，不能在放值了
	fmt.Println(d4)
	d5 := <-msg //虽然channel已经关闭，但仍可以取值，不能在放值了
	fmt.Println(d5)
	time.Sleep(time.Second)

}
