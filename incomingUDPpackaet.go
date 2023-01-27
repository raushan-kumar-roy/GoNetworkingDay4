package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	fmt.Println("Listening on 0.0.0.0:8080")

	for {
		var buf [1024]byte
		n, _, _ := conn.ReadFromUDP(buf[:])
		fmt.Println("Received:", string(buf[:n]))
	}
}
