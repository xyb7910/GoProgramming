package main

import "fmt"
import "time"

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// 尝试获取互斥锁，会创建一个计时器，并且尝试从互斥锁通道中获取值
// 如果在指定时间内读到了数据，表示获取到了互斥锁，返回 true
// 如果超时（即无法在规定时间内读取到锁），则表示无法获取到锁，返回 false

func (mu *Mutex) Lock() bool {
	timer := time.NewTimer(time.Millisecond)
	select {
	case <-mu.ch:
		timer.Stop()
		return true
	case <-timer.C:
		return false
	}
}

// 释放互斥锁，尝试向互斥锁通道中发送一个值，如果通道已满，表示互斥锁没有被 Lock

func (mu *Mutex) Unlock() {
	select {
	case mu.ch <- struct{}{}:
	default:
		panic("unlock of unheld mutex")
	}
}

func main() {
	mutex := NewMutex()

	fmt.Println("Lock the mutex...")

	if !mutex.Lock() {
		fmt.Println("Failed to lock the mutex.")
		return
	}

	fmt.Println("The mutex is locked.")

	go func() {
		fmt.Println("Try to lock the mutex in another goroutine...")
		if !mutex.Lock() {
			fmt.Println("Failed to lock the mutex in another goroutine.")
			return
		}
		fmt.Println("Never reach here.")
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("Unlock the mutex...")
	mutex.Unlock()
	time.Sleep(time.Millisecond * 50)
}
