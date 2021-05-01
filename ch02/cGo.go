package main

//#include <stdio.h>
//// test: make it four slashes
//void callC() {
//    printf("Calling C code!\n");
//}
import "C"
import "fmt"

func main() {
    // this is for test
    fmt.Println("A Go statement!")
    C.callC()
    fmt.Println("Another Go statement!")
}
