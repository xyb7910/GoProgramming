package main

import (
	"fmt"
	"time"
)

func Hu(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("after two second hu!!")

	// 执行语句，通知主协程结束
	ch <- 100
}

func Send(ch chan int) {
	for i := 0; i < 13; i++ {
		ch <- i
		fmt.Println("send ", i)
	}
	close(ch)
}

func Receive(ch chan int) {
	time.Sleep(time.Second * 2)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Printf("chan close, receicve:%v\n", v)
				return
			}
			fmt.Println("receive ", v)
		}
	}
}

func main() {
	//ch := make(chan int)
	//
	//go Hu(ch)
	//
	//fmt.Println("start hu, wait for hu finish")
	//
	//// 从缓冲信道中读取数据
	//v := <-ch
	//fmt.Println("hu finish, value is ", v)

	//ch := make(chan int, 10)
	//
	//go Receive(ch)
	//go Send(ch)
	//
	//for {
	//	time.Sleep(time.Second * 1)
	//}

	buffedChan := make(chan int, 2)
	buffedChan <- 1
	buffedChan <- 2
	close(buffedChan)
	for i := range buffedChan {
		fmt.Println(i)
	}
}
