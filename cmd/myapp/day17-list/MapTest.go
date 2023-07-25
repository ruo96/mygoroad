package main

import (
	"fmt"
)

func main() {
	// 创建一个空的 Map
	m := make(map[string]int)

	// 创建一个初始容量为 10 的 Map   make是内置函数， 主要用于初始化slice，map，channel   但是slice可以不用初始化的
	m = make(map[string]int, 10)

	fmt.Println(m)

	// 定义之后必须要初始化
	m1 := map[string]string{
		"go":   "go1",
		"grpc": "grpc1",
		"gin":  "gin1",
	}

	fmt.Println(m1["grpc"])
	m1["grpc"] = "newrpc"
	fmt.Println(m1["grpc"])

	// map是无序的，不保证每次打印的数据,map不是线程安全的
	// sync.Map{}这个是线程安全的

	delete(m1, "grpc")
	fmt.Println(m1)

}
