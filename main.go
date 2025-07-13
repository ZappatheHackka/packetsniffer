package main

import (
	"fmt"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, 1)

	if err != nil {
		fmt.Printf("Error creating socket: %v.\n", err)
	}
	defer syscall.Close(fd)

	fmt.Printf("Socket created successfully with file descriptor: %d.\n", fd)

	buffer := make([]byte, 1500)

	data, err := syscall.Read(fd, buffer) // "data" is how many bytes the packet contains
	fmt.Println("Data read.")

	if err != nil {
		fmt.Printf("Could not read data, error: %v\n", err)
	}

	fmt.Println(data)
	packet_info := buffer[0:data] // the actual packet bytes are queued in the buffer by the kernel.
	fmt.Printf("Info contained therein: %v", packet_info)
}
