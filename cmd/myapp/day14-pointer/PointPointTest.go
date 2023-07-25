package main

import "fmt"

func main() {
	var a int = 5
	// 指针
	var p1 *int = &a
	// 二级指针
	var p2 **int = &p1
	// 三级指针
	var p3 ***int = &p2

	fmt.Printf("p1的值:%d    p1的目标值:%d\n", p1, *p1)
	fmt.Printf("p2的值:%d    p2的目标值:%d    p2的链尾目标值:%d\n", p2, *p2, **p2)
	fmt.Printf("p3的值:%d    p3的目标值:%d    下一个目标值:%d    p3的链尾目标值:%d\n", p3, *p3, **p3, ***p3)

}
