package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	providerAddress := "localhost:9000"

	// Respond with the provider's address
	conn.Write([]byte(providerAddress))
}

func main() {
	fmt.Println("Orchestrator is running.")

	// Listen for incoming connections
	ln, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error starting the orchestrator:", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}
