package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SpinLock struct {
	// 1 表示 Lock
	// 0 表示 Unlock
	state int32
}

func NewSpinLock() *SpinLock {
	return &SpinLock{0}
}

// 添加上锁的逻辑

func (l *SpinLock) Lock() {
	for !l.TryLock() {

	}
}

// 解锁逻辑

func (l *SpinLock) UnLock() {
	atomic.CompareAndSwapInt32(&l.state, 1, 0)
}

// 尝试上锁逻辑，如果不能上锁的话返回 false

func (l *SpinLock) TryLock() bool {
	return atomic.CompareAndSwapInt32(&l.state, 0, 1)
}

func main() {
	s := NewSpinLock()

	go func() {
		s.Lock()
		fmt.Println("goroutine: lock success")
		s.UnLock()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("main thread: trying to get lock")
	s.Lock()
	fmt.Println("main thread: lock success")
}
