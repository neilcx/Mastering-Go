package main

// #cgo CFLAGS: -I${SRCDIR}/callClib
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Mihalis!")
	// you need a defer statement in order to free the memory space of the C string
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")

}


//steps:
//	ls -al callClib/
//	gcc -c callClib/*.c
//	ls -l callC.o
//	file callC.o
//callC.o: Mach-O 64-bit object x86_64
//	ar rs callC.a *.o
//ar: creating archive callC.a
// 	ls -l callC.a
//-rw-r--r--  1 neil  453037844  4096 May  1 15:36 callC.a
//	file callC.a
//callC.a: current ar archive random library
//	rm callC.o
//  go build callC.go
//	ls -l callC
//-rwxr-xr-x  1 neil  453037844  2594008 May  1 15:39 callC
//  file callC
////callC: Mach-O 64-bit executable x86_64
//	./callC
//Going to call a C function!
//Hello from C!
//Going to call another C function!
//Go send me This is Mihalis!
//All perfectly done!