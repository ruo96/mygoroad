package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := make(chan int)

	go add(a, b, c) // 启动一个goroutine执行add函数

	sum := <-c // 从channel中读取计算结果
	fmt.Println(sum)
}

func add(a int, b int, c chan int) {
	sum := a + b
	c <- sum // 将计算结果写入channel
}
