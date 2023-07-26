package main

import (
	"fmt"
	"time"
)

func main() {
	// 在go语言并发编程中， goroutine之间进行通信
	// 不要通过共享内存来通信，而要通过通信实现内存共享
	// php, python, java多线程编程的时候，两个线程之间通信通常都是用的全局变量
	// 也会提供消息队列的机制，消费者和生产者之间的关系，这个是go推荐的方式，也就是channel，再加上语法糖
	// 无缓冲channel，适用于通知，B要知道A是否完成；有缓冲channel，适用于消费者和生产者之间的通信
	// go中channel的应用场景， 1. 消息传递，消息过滤   2. 信号广播   3. 事件订阅和广播   4. 任务分发    5. 结果汇总    6. 并发控制    7. 同步和异步   。。。。

	var msg chan string        // 这个就是通道
	msg = make(chan string, 0) // channel的初始化值为0的话， 放值进去会阻塞，  有缓冲和无缓冲的channel具有不同的应用场景，如果设置为0就是无缓冲的channel

	//msg <- "w1"

	var a string
	a = "hello" // 对于没有缓冲区的channel而言， 先读或者先写都是deadlock
	msg <- a
	fmt.Println(a)
	//
	//fmt.Println(a)

	go func(msg chan string) { // go有一种happen before机制， 这就保证了消费这个channel是会在  输入channel之前执行，所以不会deadlock
		fmt.Println("开始接收值")
		a := <-msg
		fmt.Println("接收到值： ", a)
	}(msg)

	//time.Sleep(time.Second * 3)

	msg <- "hello world"
	fmt.Println("发送值： ", msg)

	time.Sleep(time.Second * 1)

}
