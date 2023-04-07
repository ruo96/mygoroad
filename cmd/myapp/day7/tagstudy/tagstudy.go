package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city,omitempty"`
}

func main() {
	p := Person{Name: "John Doe", Age: 30}

	b, _ := json.Marshal(p)
	fmt.Println(string(b))

	p2 := Person{Name: "Jane Doe", Age: 25, City: "New York"}

	b2, _ := json.Marshal(p2)
	fmt.Println(string(b2))
}
