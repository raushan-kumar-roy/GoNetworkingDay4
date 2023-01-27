package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server.")

	for {
		_, err := conn.Write([]byte("Hello from client 1"))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println("Received message:", string(buf))
		time.Sleep(time.Second * 5)
	}
}
