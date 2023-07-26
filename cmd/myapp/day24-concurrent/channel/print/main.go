package main

import (
	"fmt"
	"time"
)

// 交叉打印序列，一个打印数字。另外一个打印字母
// 效果： 12AB34CD

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		<-number
		// 此处应该等待另一个goroutine来通知
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printAlphaBeta() {
	i := 65
	for {
		<-letter
		// 此处应该等待另一个goroutine来通知
		fmt.Printf("%c%c", i, i+1)
		if fmt.Sprintf("%c", i+1) == "Z" {
			return
		}
		i += 2
		number <- true
	}

}

func main() {

	go printNum()
	go printAlphaBeta()

	//fmt.Printf("%d%d", num.Add(1), num.Add(1))
	number <- true

	time.Sleep(time.Second * 1)

	fmt.Println(string(rune(65)))
	fmt.Println(string(65))

}
