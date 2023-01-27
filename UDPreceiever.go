package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	fmt.Println("Sending to", addr.String())
	for {
		_, err := conn.Write([]byte("Hello, Server!"))
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
