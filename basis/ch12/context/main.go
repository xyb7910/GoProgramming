package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//尽量不要使用共享变量
func cpuInfo(ctx context.Context) {
	//希望拿到一个请求的id
	fmt.Printf("traceid: %s\r\n", ctx.Value("traceid"))

	//记录一些日志，这次请求是哪个traceid打印的

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu信息")
		}
	}
}

func main() {
	//var stop = make(chan struct{})
	wg.Add(1)
	//context包中包含了三个函数,如果希望被控制， 超时，传值，但是我不希望影响我原来接口信息的时候，函数参数中第一个参数尽量加上ctx

	//ctx, cannel := context.WithCancel(context.Background())
	//ctx1, _ := context.WithCancel(ctx)

	//timeout 设置主动超时机制
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//go cpuInfo(ctx) //调用具有传递性

	//WithDeadLine 在时间点上cannel

	//WithValue 通过值来进行cannel
	valueCtx := context.WithValue(ctx, "traceid", "ypb")
	go cpuInfo(valueCtx)
	//time.Sleep(6 * time.Second)
	//cannel()
	//stop <- struct{}{}
	wg.Wait()
	fmt.Println("监控完成")
}
