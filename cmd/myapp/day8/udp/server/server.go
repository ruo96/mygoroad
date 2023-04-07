package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error listening on UDP:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening for UDP messages on port 8080...")

	for {
		buf := make([]byte, 1024)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving message:", err)
			continue
		}

		fmt.Println("Received message from", clientAddr.String(), ":", string(buf[:n]))

		_, err = conn.WriteToUDP([]byte("Hello, client!"), clientAddr)
		if err != nil {
			fmt.Println("Error sending response message:", err)
			continue
		}
	}
}
