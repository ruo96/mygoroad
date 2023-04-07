package day2

import "fmt"

func CalIota() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
		h = iota
		i
	)

	const (
		i1 = 1 << iota
		i2 = 3 << iota
		i3
		i4
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
	fmt.Println(i1, i2, i3, i4)

	a1 := 1
	fmt.Println(&a1)
	p := &a1
	fmt.Println(*p)
}
