package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁本质上是把并行的串行化
// 即使是设计锁，那么也应该尽量保证并行
// 我们有两组协程，一组负责写数据，一组负责读数据，web系统中绝大部分场景都是读多写少
// 虽然有多个goroutine，但是仔细分析我们会发现，读协程之间应该并发，读和写之间应该互斥串行，写写之间也是要互斥串行
var rw sync.RWMutex
var total int

func write() {
	rw.Lock()
	total += 1
	rw.Unlock()
}

func read() {
	rw.RLocker().Lock()
	total += 1
	rw.RLocker().Unlock()
}

func main() {
	//var num int
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)

	// 写的goroutine
	go func() {
		defer wg.Done()
		rwlock.Lock() // 加写锁，会防止别的写锁获取和读锁获取
		defer rwlock.Unlock()
		fmt.Println("get write lock")
		time.Sleep(5 * time.Second)

	}()

	time.Sleep(time.Second)

	go func() {
		defer wg.Done()

		for {
			rwlock.RLock() // 加了读锁，不会阻止别人来读取
			time.Sleep(500 * time.Millisecond)
			fmt.Println("get read lock")
			rwlock.RUnlock()
		}

	}()

	wg.Wait()
}
