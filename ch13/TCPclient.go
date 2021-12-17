package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	// Examples:
	//	Dial("tcp", "golang.org:http")
	//	Dial("tcp", "192.0.2.1:http")
	//	Dial("tcp", "198.51.100.1:80")
	//	Dial("udp", "[2001:db8::1]:domain")
	//	Dial("udp", "[fe80::1%lo0]:53")
	//	Dial("tcp", ":80")
	//

	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
