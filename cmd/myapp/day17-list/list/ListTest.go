package main

import (
	"container/list"
	"fmt"
)

func main() {
	//myList := list.New()
	var myList list.List
	myList.PushBack(1)
	myList.PushBack("w1")
	myList.PushBack(2)
	myList.PushBack("w2")
	myList.PushFront(0)

	fmt.Println(myList)
	fmt.Println(myList.Front().Value.(string))

}
