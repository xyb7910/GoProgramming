package main

import "fmt"

//func main() {
//	chan1 := make(chan int)
//	go rev(chan1)
//	chan1 <- 1
//	fmt.Println("send success")
//}

func rev(c chan int) {
	ret := <-c
	fmt.Println("receive success", ret)
}

// 创建两个channel,将1到100写进chan1中， 并且将chan1中的数据取出乘以二写进chan2中
func main() {
	//chan1 := make(chan int, 100)
	//chan2 := make(chan int, 100)
	//
	//go send(chan1)
	//go receive(chan1, chan2)
	//
	//for res := range chan2 {
	//	fmt.Println(res)
	//}
	//defer close(chan2)
	//defer close(chan1)

	// 打印十以内的奇数
	ch := make(chan int, 1)
	for i := 1; i < 10; i++ {
		select {
		case x := <-ch:
			{
				fmt.Println(x)
			}
		case ch <- i:
			{

			}
		}
	}
}

func send(c chan int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
}

func receive(c1 chan int, c2 chan int) {
	for {
		num, ok := <-c1
		if ok {
			num = num * 2
			c2 <- num
		} else {
			break
		}
	}
}
