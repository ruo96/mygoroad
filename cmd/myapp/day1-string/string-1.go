package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "这里是 www\n.runoob\n.com"
	fmt.Println("-------- 原字符串 ----------")
	fmt.Println(str)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println("-------- 去除空格与换行后 ----------")
	fmt.Println(str)

	a := "hello"
	b := "hello"

	fmt.Println(a >= b)

	compare := strings.Compare(a, b)
	fmt.Println("compare result is : ", compare)

	s := "go go d 世界 go"
	count := strings.Count(s, "go")
	fmt.Println(count)

	split := strings.Split(s, " ")
	fmt.Println(split)

	prefix := strings.HasPrefix(s, "g")
	fmt.Println(prefix)

	suffix := strings.HasSuffix(s, "o")
	fmt.Println(suffix)

	index := strings.Index(s, "go")
	fmt.Println(index)

	indexRune := strings.IndexRune(s, '世')
	fmt.Println(indexRune)

	replace := strings.Replace(s, "go", "do", -1) //最后的数字代表替换多少个， 负数代表全部替换
	fmt.Println(replace)

	lower := strings.ToLower("ABC")
	fmt.Println(lower)

	upper := strings.ToUpper(lower)
	fmt.Println(upper)

	trim := strings.Trim(s, " ")
	fmt.Println(trim)

	num := 2

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}
}
