package main

import (
	"fmt"
	"sync"
)

func main() {
	// 连接数据库，打开文件，开始锁，无论如何，最后都要记得去关闭数据库，关闭文件，解锁
	// 类似于java里面的finally
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	
	defer fmt.Println("defer 1") // 这个会后执行
	defer fmt.Println("defer 2") // 这个会先执行

	fmt.Println("noraml 1")
	fmt.Println("noraml 2")
}
