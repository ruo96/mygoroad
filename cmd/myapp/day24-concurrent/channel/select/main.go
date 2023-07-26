package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{}) // channel是多线程安全的 channel要初始化
// var done chan struct{} // channel是多线程安全的
func g1(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{} // 这个是比较常用的写法
}

func g2(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{} // 这个是比较常用的写法
}

func main() {
	// select 类似于switch case语句， 但是select的功能和操作linux里面提供的io的select，poll，epoll差不多
	// select主要作用于多个channel，和其密不可分
	// 当两个goroutine在执行的时候，在主gouroutine中， 当某一个协程完成之后，需要立马知道 ,有一个方法是通过共享变量， 但是go并不推崇
	g1Channel := make(chan struct{})
	g2Channel := make(chan struct{})
	go g1(g1Channel)
	go g2(g2Channel)

	// 需要监控多个channel， 任何一个channel返回都可以知道
	//<-done // 这个是阻塞的
	// 某一个分支就绪了，就执行该分支，如果两个都就绪了，先执行哪个？ 这个是随机的，目的是：防止饥饿，
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("timeout")
			return
		}
	}

	fmt.Println("done")
}
