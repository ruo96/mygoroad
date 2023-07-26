package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("consumer: ", num)
	}
}

func main() {

	// 需要单向的channel， 让一方只接收数据

	//var ch1 chan int
	//var ch2 chan<- float64 // 这种就是单向的channel， 只能写入float64
	//var ch3 <-chan int     // 单向的，只能读取
	//
	//a := <-ch3
	//fmt.Println(a)

	//c := make(chan int, 3)
	//var send chan<- int = c // send only
	//
	//send <- 1
	//
	//var read <-chan int = c
	//
	//var a = <-read
	//fmt.Println(a)

	//d1 := (chan int)send  不能够再转换回去

	c := make(chan int)
	go producer(c)
	go consumer(c)

	time.Sleep(time.Second * 10)

}
