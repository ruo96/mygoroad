package main

import "fmt"

type Person struct {
	name string
	age  int
}

func change(p Person) {
	p.name = "changed"
}

func change2(p *Person) {
	p.name = "changed"
}

func main() {
	person := Person{
		name: "w1",
	}

	change2(&person)

	fmt.Println(person)

	//	指针的初始化
	ps := &Person{} // 第一种初始化方式

	fmt.Println("1---", ps.age)

	var emptyPerson Person // 第二种初始化方式
	pi := &emptyPerson
	fmt.Println("2---", pi.age)

	var pp *Person = new(Person) // 第三种初始化方式
	fmt.Println("3---", pp.age)

	// 初始化两个关键字， map，channel，slice初始化推荐使用make方法
	// 指针初始化推荐使用new函数, 指针需要进行初始化， 否则会出现nil pointer
	// map必须初始化

	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, "-------------", b)

	c := &a
	d := &b
	*c, *d = *d, *c
	fmt.Println(a, "-------------", b)

}
