package main

import "C"

import (
	"fmt"
)

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {
}


// go build -o usedByC.o -buildmode=c-shared usedByC.go
//ls -al usedByC.*
//-rw-r--r--  1 chenxi1  453037844      263 May  1 18:51 usedByC.go
//-rw-r--r--  1 chenxi1  453037844     1639 May  1 18:51 usedByC.h
//-rw-r--r--  1 chenxi1  453037844  2599544 May  1 18:51 usedByC.o
//file usedByC.o
//usedByC.o: Mach-O 64-bit dynamically linked shared library x86_64