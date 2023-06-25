package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		randomNumber := rand.Intn(100) + 1
		currentTime := time.Now().Unix()

		response := fmt.Sprintf("Time: %d, Random Number: %d\n", currentTime, randomNumber)
		conn.Write([]byte(response))

		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("Service Provider is running.")

	// Listen for incoming connections
	ln, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error starting the service provider:", err)
		os.Exit(1)
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
