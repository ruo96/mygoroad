package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cpuInfo(ctx context.Context) {
	fmt.Printf("key : %s\r\n", ctx.Value("key1"))
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println(" cpu info ")
		}
	}
}

func main() {
	// 渐进式的方式
	// 有一个goroutine来监控cpu的信息
	wg.Add(1)

	// context提供了三种函数， withCancel withTimeout withValue
	//ctx, cancel := context.WithCancel(context.Background()) // 具有传递性
	//ctx2, _ := context.WithCancel(ctx)

	// 主动超时 timeout  还有withdeadline
	ctx, _ := context.WithTimeout(context.Background(), time.Second*6)

	// withvalue
	ctx1 := context.WithValue(ctx, "key1", "value")
	fmt.Println(ctx1)

	go cpuInfo(ctx1)

	//time.Sleep(time.Second * 6)
	//cancel() // 父的cancel方法也都可以控制
	wg.Wait()
	fmt.Println("监控完成")
}
