package main

import "fmt"

func main1() {
	a := 10
	fmt.Println(&a)

	b := [5]int{1, 2, 3, 4, 5}
	var ptr [5]*int
	var i int
	for i = 0; i < 5; i++ {
		ptr[i] = &b[i]
	}
	for i = 0; i < 5; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

}
