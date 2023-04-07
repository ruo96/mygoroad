package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string //Name字段首字母大写
	Age  int    //age字段首字母小写          大写的变量才可以序列化， 小些的是不会json装换的  定义的结构体如果只在当前包内使用，结构体的属性不用区分大小写。如果想要被其他的包引用，那么结构体的属性的首字母需要大写。例如
}

func main() {
	person := Person{"小明", 18}
	if result, err := json.Marshal(&person); err == nil { //json.Marshal 将对象转换为json字符串
		fmt.Println(string(result))
	}
}
