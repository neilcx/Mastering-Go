package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {

	//	AF_INET                           = 0x2
	// 	SOCK_RAW                          = 0x3
	// 	IPPROTO_ICMP                      = 0x1
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		fmt.Println("Error in syscall.Socket:", err)
		return
	}

	f := os.NewFile(uintptr(fd), "captureICMP")
	if f == nil {
		fmt.Println("Error in os.NewFile:", err)
		return
	}

	err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 256)
	if err != nil {
		fmt.Println("Error in syscall.Socket:", err)
		return
	}

	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("% X\n", buf[:numRead])
	}
}
