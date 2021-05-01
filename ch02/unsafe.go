package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)

	//➜  ch02 git:(master) ✗ ./unsafe
	//*p1:  5
	//*p2:  5
	//5434123412312431212
	//*p2:  -930866580
	//54341234
	//*p2:  54341234
}
