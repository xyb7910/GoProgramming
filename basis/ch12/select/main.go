package main

import (
	"fmt"
	"time"
)

func g1(ch chan struct{}) {
	time.Sleep(time.Second)
	ch <- struct{}{}
}

func g2(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}

func main() {
	//select 类似于 switch case语句，但是select的功能和我们操作linux里面提供的io的select, poll, epoll
	//select可作用于多个channel

	//现在有个需求，现在有两个goroutine都在执行，但是我在住的goroutine中，当某一个执行完成以后，这个时候我会立马知道

	g1channel := make(chan struct{})
	g2channel := make(chan struct{})

	//g1channel <- struct{}{}
	//g2channel <- struct{}{}

	go g1(g1channel)
	go g2(g2channel)

	//监控多个channel，任何一个channel返回都要知道
	//1、某个分支就绪了就执行该分支 2、如果两个都就绪了，执行的顺序是随机的， 目的是防止饥饿

	timer := time.NewTimer(5 * time.Second) //超时
	for {
		select {
		case <-g1channel:
			fmt.Println("g1 done")
		case <-g2channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("time out")
			return
		}
	}

}
