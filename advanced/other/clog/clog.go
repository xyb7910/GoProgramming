package clog

import (
	"fmt"
	"sync"
	"time"
)

var chokeTime = time.Second * 1

// sync.WaitGroup

func chokeWithWaitGroup() {
	start := time.Now()

	wg := sync.WaitGroup{}

	// 添加阻塞计时器
	wg.Add(1)
	go func() {
		time.Sleep(chokeTime)
		wg.Done()
	}()
	wg.Wait()
	d := time.Since(start)
	fmt.Println("使用WaitGroup阻塞了:", d)
}

// select

func chokeWithSelect() {
	start := time.Now()

	// 启动定时器
	timer := time.NewTimer(chokeTime)
	select {
	case <-timer.C:
		d := time.Since(start)
		fmt.Println("使用Select阻塞了:", d)
	}
}

// channel

func chokeWithChannel() {
	start := time.Now()
	ch := make(chan struct{})
	go func() {
		time.Sleep(chokeTime)
		ch <- struct{}{}
	}()
	// 阻塞，等待 channel 数据
	if _, ok := <-ch; ok {
		d := time.Since(start)
		fmt.Println("使用channel阻塞了:", d)
	}
}

// time.After

func chokeWithTimeAfter() {
	start := time.Now()
	<-time.After(chokeTime)
	d := time.Since(start)
	fmt.Println("使用time.After阻塞了:", d)
}

// for

func chokeWithFor() {
	start := time.Now()
	for {
		d := time.Since(start)
		if d > chokeTime {
			fmt.Println("使用 For 阻塞了:", d)
			return
		}
	}
}

// mutex

func chokeWithMutex() {
	start := time.Now()
	mu := sync.Mutex{}

	mu.Lock()

	go func() {
		mu.Lock()
		defer mu.Unlock()

		d := time.Since(start)
		fmt.Println("使用 mutex 阻塞了:", d)
	}()
	time.Sleep(chokeTime)
	mu.Unlock()
	time.Sleep(chokeTime)
}
