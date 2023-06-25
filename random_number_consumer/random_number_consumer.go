package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("Service Consumer is running.")

	// Connect to the orchestrator
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting to the orchestrator:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Read the provider's address
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading provider's address:", err)
		os.Exit(1)
	}

	providerAddress := string(buffer[:n])

	for {
		// Connect to the provider
		conn, err := net.Dial("tcp", providerAddress)
		if err != nil {
			fmt.Println("Error connecting to the service provider:", err)
			continue
		}
		defer conn.Close()

		// Read and print the response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response from the service provider:", err)
			continue
		}

		fmt.Print(string(buffer[:n]))

		time.Sleep(1 * time.Second)
	}
}
