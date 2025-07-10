package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 1)

	if err != nil {
		fmt.Printf("Error creating socket: %v.\n", err)
	}
	defer syscall.Close(fd)

	fmt.Printf("Socket created successfully with file descriptor: %d.\n", fd)
}
