package main

import (
	"fmt"
)

// panic会导致程序的退出，平时开发中不要随便用， 一般在一个服务启动过程中，如果有依赖没有准备好， 那么可以主动调用panic；但是服务一旦启动了， 一旦不小心panic，那么整个程序就会挂掉
// 但是有的时候是不小心导致被动触发panic
// defer需要放在panic之前定义，另外recover只有在defer调用的函数中才会生效
// recover处理异常后，逻辑并不会恢复到panic那个点去
func main() {
	fmt.Println("外层开始")
	defer func() {
		fmt.Println("外层准备recover")
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "外层", err) // err已经在上一级的函数中捕获了，这里没有异常，只是例行先执行defer，然后执行后面的代码
		} else {
			fmt.Println("外层没做啥事")
		}
		fmt.Println("外层完成recover")
	}()
	fmt.Println("外层即将异常")
	f()
	fmt.Println("外层异常后")
	defer func() {
		fmt.Println("外层异常后defer")
	}()
}

func f() {
	fmt.Println("内层开始")
	defer func() {
		fmt.Println("内层recover前的defer")
	}()

	defer func() {
		fmt.Println("内层准备recover")
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "内层", err) // 这里err就是panic传入的内容
		}

		fmt.Println("内层完成recover")
	}()

	defer func() {
		fmt.Println("内层异常前recover后的defer")
	}()

	panic("异常信息")

	defer func() {
		fmt.Println("内层异常后的defer")
	}()

	fmt.Println("内层异常后语句") //recover捕获的一级或者完全不捕获这里开始下面代码不会再执行
}
