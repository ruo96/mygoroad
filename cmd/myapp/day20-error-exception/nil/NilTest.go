package main

import "fmt"

type Person struct {
	name string
	age  int
	f    *int
}

func main() {
	// 不同类型的数据，零值是不一样的
	//bool false;
	// number 0;   string "", pointer nil,   slice nil,  map nil, channel/interface/function nil
	// struct  默认值不是nil， 默认值是具体字段的默认值

	p1 := Person{
		name: "w1", age: 10,
	}

	p2 := Person{
		age: 10, name: "w1",
	}

	if p1 == p2 {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}

	var ps []Person //nil slice
	if ps != nil {
		fmt.Println("slice not nil")
	} else {
		fmt.Println("slice is nil")
	}

	var ps2 = make([]Person, 0) // empty slice
	if ps2 != nil {
		fmt.Println("slice2 not nil")
	} else {
		fmt.Println("slice2 is nil")
	}
}
