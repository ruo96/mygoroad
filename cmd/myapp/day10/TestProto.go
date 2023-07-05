package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	zetyun_person "mygo/cmd/myapp/day10/zetyun.person"
)

func main() {
	p := &zetyun_person.User{
		Name:  "John Doe",
		Age:   42,
		Email: "jdoe@example.com",
	}

	data, err := proto.Marshal(p)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	fmt.Println("Marshalled data:", data)
	fmt.Println("\n")

	var p2 zetyun_person.User
	err = proto.Unmarshal(data, &p2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p2)

}
