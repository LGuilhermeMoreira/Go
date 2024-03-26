package main

import (
	"flag" // Package flag implements command-line flag parsing.
	"fmt"  // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"net"  // Package net provides a portable interface for network I/O, including TCP/IP and UDP.
	// read: https://okanexe.medium.com/the-complete-guide-to-tcp-ip-connections-in-golang-1216dae27b5a
	// read: https://pkg.go.dev/net#pkg-overview
	// These comments provide references for further reading about TCP/IP connections in Go and the net package.
)

func main() {
	// Define a variable to store the IP address.
	var ip string

	// Define a flag named "ip" that represents the IP address.
	flag.StringVar(&ip, "ip", "", "IP address to connect to")

	// Parse the command-line flags.
	flag.Parse()

	// Establish a TCP connection to the specified IP address.
	conn, err := net.Dial("tcp", ip) // Protocol and IP/IP+PORT/PORT
	if err != nil {
		// "panic" stops the normal execution of the program and prints the error message.
		panic(err)
	}

	// Defer the closing of the connection until the function returns.
	defer conn.Close()

	// Create a buffer to store data received from the server.
	buffer := make([]byte, 1024)

	// Read data from the connection into the buffer.
	bytesRead, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	// Print the remote address (IP and port) of the server and the data received from the server.
	fmt.Printf("Remote address: %v\nData from server: %v", conn.RemoteAddr(), string(buffer[:bytesRead]))
}
