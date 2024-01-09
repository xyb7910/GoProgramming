package main

import (
	"fmt"
	"time"
)

// 两个goroutine之间通信最常用的方式是 一个全局变量
func main() {
	var msg chan string
	// 无缓冲channel适用于 通知， B要第一时间知道A是否已经完成
	// 有缓冲channel适用于消费者和生产者之间的通信
	/*
		go中的channel的场景：
		消息传递，消息过滤
		消息广播
		事件订阅 和 广播
		任务分发
		结果汇总
		并发控制
		同步和异步
	*/
	//msg = make(chan string, 1) //channel 初始化值为0的话，放值的时候会被阻塞
	msg = make(chan string, 0)

	//消费无缓冲型channel， go语言中有一种默认的happen-before机制
	go func(msg chan string) {
		data := <-msg
		fmt.Println(data)
	}(msg)
	msg <- "ypb" // 放值到channel中

	//waitgroup 如果少了done调用，容易出现deadlock， 无缓冲的channel在放值和读取值顺序不当的时候也容易出现deadlock
	time.Sleep(10 * time.Second)

}
