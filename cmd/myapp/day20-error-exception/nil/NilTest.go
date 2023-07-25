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

	var m map[string]string
	if m != nil {
		fmt.Println("m not nil")
	} else {
		fmt.Println("m is nil")
	}

	for key, value := range m {
		fmt.Println(key, value) // 不会抛出异常
	}

	var m2 = make(map[string]string, 0)
	if m2 != nil {
		fmt.Println("m2 not nil")
	} else {
		fmt.Println("m2 is nil")
	}

	for key, value := range m2 {
		fmt.Println(key, value) // 不会抛出异常
	}

	fmt.Println(m["name"])
	m["name"] = "w1" // 尽管打印key，value等操作不会报错，但是这个赋值却会报错 panic， 因此最好都用make进行初始化操作
}
