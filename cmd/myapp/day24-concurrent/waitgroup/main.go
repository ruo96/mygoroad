package main

import (
	"fmt"
	"sync"
)

func asyncPrint() {
	//time.Sleep(time.Duration(1000))
	fmt.Println("hello")
}

// 主协程
func main() {

	var wg sync.WaitGroup

	// 我要监控多少个goroutine执行结束
	wg.Add(100)

	for i := 0; i < 100; i++ {
		// 也可以在这边增加一个wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
