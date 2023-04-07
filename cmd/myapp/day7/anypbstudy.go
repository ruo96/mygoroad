package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/types/known/anypb"
)

type MyData struct {
	Name string
	Age  int32
}

func main() {
	myData := MyData{Name: "John", Age: 25}

	// Marshal MyData to Any.
	any, err := anypb.New(proto.MessageV2(&myData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Unmarshal Any to MyData.
	var unmarshaled MyData
	if err := any.UnmarshalTo(&unmarshaled); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Compare original MyData and unmarshaled MyData using go-cmp.
	if diff := cmp.Diff(myData, unmarshaled); diff != "" {
		fmt.Println("Error: mismatch:", diff)
		return
	}

	// Marshal Any to JSON.
	json, err := any.MarshalJSON()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(json))
}
