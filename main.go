package main

import (
	"fmt"
	"net"
)

func is_port_open(ip string, port int) bool {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	ip := "127.0.0.1"
	fmt.Println("Scanning for open ports..")
	for port := 1; port <= 65535; port++ {
		if is_port_open(ip, port) {
			fmt.Printf("Port %d is open\n", port)
		}
	}

	fmt.Println("Scan complete.")
}
