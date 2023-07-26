package main

import "fmt"

const name = 1
const (
	v1 = iota
	v2
	v3
	v4 = "w1"
	v5
	v6 = 7
	v7 = iota
)

func main() {
	fmt.Println("hello")
}

func add(a, b int) int {
	return a + b
}
