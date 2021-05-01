package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	myVersion := runtime.Version()
	fmt.Println(myVersion)
	major := strings.Split(myVersion, ".")[0][2] //gox, x is in position [2]
	fmt.Println( string(major))
	minor := strings.Split(myVersion, ".")[1]
	fmt.Println(minor)
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(minor)

	if m1 == 1 && m2 < 8 {
		fmt.Println("Need Go version 1.8 or higher!")
		return
	}

	fmt.Println("You are using Go version 1.8 or higher!")
}
