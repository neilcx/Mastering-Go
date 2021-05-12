package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	fmt.Println(x)
}


// still print x once
func writeToChannel2(c2 chan int, x int) {
	fmt.Println(x)
	c2 <- x
	fmt.Println("haha")
	fmt.Println(x)
	close(c2)
}


func main() {
	c := make(chan int)
	c2 := make(chan int)
	go writeToChannel(c, 10)

	time.Sleep(2 * time.Second)

	go writeToChannel2(c2, 20)

	time.Sleep(2 * time.Second)
}


//go run writeCh.go 10
//10
//20
