package main

import "fmt"

func main1() {
	add_func := add(1, 2)
	fmt.Println(add_func())
	fmt.Println(add_func())
	fmt.Println(add_func())
}

// 闭包使用方法  函数里面包含匿名函数，然后操作原函数内部的变量，就是闭包
func add(x1, x2 int) func() (int, int) {
	i := 0
	return func() (int, int) {
		i++
		return i, x1 + x2
	}
}
