package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.2:8080")
	if err != nil {
		fmt.Println("Error dialing server:", err)
		return
	}

	defer conn.Close()

	for {
		fmt.Fprintf(conn, "Log information: %v\n", time.Now())
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Response: " + response)
		time.Sleep(5 * time.Minute)
	}
}
