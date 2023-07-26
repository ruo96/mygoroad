package main

import "fmt"

type MyWriter interface {
	Writer(string)
}

type MyReader interface {
	Read() string
}

type MyReaderWriter interface {
	MyWriter
	MyReader
	ReadWrite()
}

type SreadWriter struct {
}

func (s *SreadWriter) Writer(s2 string) { //这个地方非常的重要
	//TODO implement me
	fmt.Println("write")
}

func (s *SreadWriter) Read() string {
	//TODO implement me
	fmt.Println("read")
	return ""
}

func (s *SreadWriter) ReadWrite() {
	//TODO implement me
	fmt.Println("read write")
}

func main() {
	var mrw MyReaderWriter = &SreadWriter{} // 如果实现的时候绑定的是值对象的方式，那么这边可以用 加&和不加两种都可以，但是如果实现的时候是用的指针对象的方式， 那么这里就只能用&的方式
	mrw.Read()
}
