package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	//time.Sleep(time.Duration(1000))
	fmt.Println("hello")
}

// 主协程
func main() {
	fmt.Println("begin")
	go asyncPrint()

	time.Sleep(1 * time.Second)
	fmt.Println("end")

	go func() {
		fmt.Println("anonymous call")
	}()

	time.Sleep(1 * time.Second)
}
