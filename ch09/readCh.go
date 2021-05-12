package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println("1", x)
	fmt.Println("tag 1")
	c <- x
	fmt.Println("tag 2")
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	time.Sleep(1 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

//go run readCh.go
//1 10
//tag 1
//Read: 10
//tag 2
//2 10
//Channel is closed!
