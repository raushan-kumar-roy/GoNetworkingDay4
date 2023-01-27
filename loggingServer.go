package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()

	fmt.Println("Listening on 0.0.0.0:8080")

	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		logInfo, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Log info: " + logInfo)
		fmt.Fprintf(conn, "Log received\n")
	}
}
