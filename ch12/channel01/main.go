package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Printf("num= %d\r\n", num)
	}
}

func main() {
	//默认情况下 channel是双向的
	// 但是在平时，我们可能需要使用单向的channel作为参数进行传递
	/*
		var ch1 chan int // 双向
		var ch2 chan<- int //单向channel，只能写入int型
		var ch3 <-chan int //单向channel，只能读取
	*/

	/*
		c := make(chan int, 3)
		var send chan<- int = c //send-only
		var read <-chan int = c //recv-only

		send <- 1 // 传入值
		<- read //读取值
	*/
	c := make(chan int)
	go producer(c)
	go consumer(c)

	time.Sleep(10 * time.Second)
}
