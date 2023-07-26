package main

import (
	"fmt"
	"strconv"
)

func add(a, b interface{}) interface{} {
	//ai, _ := a.(int)
	//bi, _ := b.(int)
	//return ai + bi

	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("not support")
	}
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))

	intsize := strconv.IntSize
	fmt.Println(intsize)
}
