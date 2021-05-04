package main

import "fmt"

func function1(i int) int {
	fmt.Printf("function 1 %d\n", i + 1) // step 3, step 5
	return i + i    // hahaha that is funny, I thought it was i + 1 ......
}

func function2(i int) int {
	return i * i
}

func funFun(f func(int) int, v int) int {
	fmt.Printf("funFun %d\n", v) // step 1
	x:=f(v) // step 2
	//y:= function1(v) // step 4
	fmt.Printf("value x: %d\n", x)
	//fmt.Printf("value y: %d\n", y)
	return x
}

func main() {
	fmt.Println("function1:", funFun(function1, 10))
	fmt.Println("===")
	fmt.Println("function2:", funFun(function2, 10))
	fmt.Println("===")
	fmt.Println("Inline", funFun(func(i int) int { return i * i * i }, 10))
}
