package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello world")

	inta := "123"
	a, _ := strconv.Atoi(inta)

	fmt.Println(a)

	str := "3.16"
	num, _ := strconv.ParseFloat(str, 64)
	fmt.Println(num)

	strNum := strconv.FormatFloat(num, 'f', 1, 64) // 会四舍五入
	fmt.Println(strNum)

	//  接口类型转换有两种情况：类型断言和类型转换。
	//
	//类型断言用于将接口类型转换为指定类型
	var i interface{} = "Hello, World"
	str1, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str1)
	} else {
		fmt.Println("conversion failed")
	}

	// 类型转换用于将一个接口类型的值转换为另一个接口类型
	var w Writer = &StringWriter{}
	sw := w.(*StringWriter)
	sw.str = "Hello, World"
	fmt.Println(sw.str)
}

type Writer interface {
	Write([]byte) (int, error)
}

type StringWriter struct {
	str string
}

func (sw *StringWriter) Write(data []byte) (int, error) {
	sw.str += string(data)
	return len(data), nil
}
