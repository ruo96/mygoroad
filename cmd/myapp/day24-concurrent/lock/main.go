package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 互斥锁,解决资源竞争

var total int32
var wg sync.WaitGroup

//var lock sync.Mutex // 锁使用的话不能复制， 否则使用起来的话，就无效， 除非就是想要两把不同的锁，就是说复制之后就是不同的锁了

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, 1)
		//lock.Lock()
		//total += 1
		//lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, -1)
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)

}
