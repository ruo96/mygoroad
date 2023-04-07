package main

import (
	"fmt"
	"net"
	"sync"
)

const bufferSize = 1024

func main() {
	// 创建一个UDP地址
	addr, err := net.ResolveUDPAddr("udp", ":8888")
	if err != nil {
		panic(err)
	}

	// 创建一个UDP连接
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	// 创建goroutine池
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte, bufferSize)
		},
	}

	// 启动多个goroutine处理客户端请求
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// 从池中获取buffer
				buf := pool.Get().([]byte)

				// 读取客户端发送的数据
				n, addr, err := conn.ReadFromUDP(buf)
				if err != nil {
					fmt.Println("Error reading from UDP connection:", err)
				}

				// 处理客户端请求
				fmt.Printf("Received %d bytes from %s: %s\n", n, addr, string(buf[:n]))

				// 将buffer归还到池中
				pool.Put(buf)
			}
		}()
	}

	// 使用select处理多个客户端连接
	for {
		select {
		case <-make(chan struct{}):
		}
	}

	//	在该示例中，我们首先创建了一个UDP地址和连接，然后创建了一个大小为10的goroutine池用于处理客户端请求。
	//	每个goroutine从池中获取buffer，并读取客户端发送的数据，然后进行处理，并将buffer归还到池中。
	// 	同时，我们使用select在一个goroutine中处理多个客户端连接，减少了goroutine的切换和创建次数，提高了服务端的性能。
}
