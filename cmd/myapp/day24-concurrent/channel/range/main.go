package main

import (
	"fmt"
	"time"
)

func main() {

	var msg chan int
	msg = make(chan int, 2)

	go func(msg chan int) {
		for a := range msg {
			fmt.Println(a)
		}

		fmt.Println("all done")
	}(msg)

	msg <- 1
	msg <- 2
	close(msg) // 而且已经关闭的channel就不能再放值了， 但是可以取值，

	data := <-msg
	fmt.Println(data)

	time.Sleep(time.Second * 3)

}
