package main

import (
	"fmt"
	"time"
)

// 声明 mymutex 结构体

type MyMutex chan struct{}

func NewMyMutex() *MyMutex {
	ch := make(chan struct{}, 1)
	mutex := MyMutex(ch)
	return &mutex
}

// 加锁时，向 channel 塞一个数据
// 如果已经被加锁，后面的数据无法加锁，陷入阻塞态

func (m *MyMutex) Lock() bool {
	select {
	case *m <- struct{}{}:
		return true
	default:
		return false
	}
}

// 解锁的时候从 channel 中取出一个数据即可

func (m *MyMutex) UnLock() {
	select {
	case <-*m:
	default:
		panic("Unlock of unlock mutex")
	}
}

func main() {
	mutex := NewMyMutex()

	if !mutex.Lock() {
		fmt.Println("unable to mutex")
		return
	}

	fmt.Println("locked mutex")

	go func() {
		if !mutex.Lock() {
			fmt.Println("unable to lock mutex in goroutine A")
			return
		}
		fmt.Println("locked mutex in goroutine A")
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("unlocking mutex...")
	mutex.UnLock()

	time.Sleep(time.Second * 2)
}
