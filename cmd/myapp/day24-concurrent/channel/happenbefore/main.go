package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // 发送操作
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // 关闭操作
}

func main() {
	ch := make(chan int)

	go sender(ch)

	for num := range ch { // 接收操作
		fmt.Println("Received:", num)
	}
}
