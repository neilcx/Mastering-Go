package main

import (
	"fmt"
	"time"
)

func function() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()
}

func main() {
	go function()

	go func() {
		for i := 10; i < 20; i++ {
			fmt.Print(i, " ")
		}
	}()

	time.Sleep(1 * time.Second)
}




//go run simple.go
//0123410 11 12 13 14 56789
//15 16 17 18 19
//
//
//go run simple.go
//
//10 11 12 13 14 15 16 17 18 19 0123456789
//
//
//go run simple.go
//0123456789
//10 11 12 13 14 15 16 17 18 19                                                                                                                                                                          ➜  ch09 git:(master)