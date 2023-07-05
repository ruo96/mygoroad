package main

import "fmt"

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber()) //这个执行结果是1
	fmt.Println(nextNumber()) //这个执行结果是2
	fmt.Println(nextNumber()) //这个执行结果是3

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence() //当getSequence()被重新赋值之后，nextNumber的值应该销毁丢失的，但并没有
	fmt.Println(nextNumber1())   //这儿因为是新赋值的，所以是1

	fmt.Println(nextNumber())  //这一行代码是补充上例子的。这儿可不是新赋的值，重点说明这一个，这儿执行居然是4，这个值并没有被销毁，原因就是闭包导致的，尽管外面的函数销毁了，但是内部函数仍然存在，还可以继续走。这个就是闭包
	fmt.Println(nextNumber1()) //新赋值的，继续执行是2
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func getRoom() int {
	var a []int32
	a[0] = 1
	b := [10]float32{1, 2, 3, 4}
	b[1] = 3
	c := [...]float32{2, 3, 4, 5}
	c[1] = 3
	d := [5]float32{1: 2, 3: 7}
	d[2] = 4
	return 0
}
