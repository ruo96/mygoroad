package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}

	num = 123
	str = strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num, str)*/

	/*str := "3.14"
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
	}*/

	num := 3.14
	str := strconv.FormatFloat(num, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num, str)
}
