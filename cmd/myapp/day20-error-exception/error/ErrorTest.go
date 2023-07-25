package main

import (
	"errors"
	"fmt"
)

func main() {
	_, err := A()
	if err != nil {
		fmt.Println(err)
	}
}

func A() (int, error) {
	return 0, errors.New("this is an error")
}
