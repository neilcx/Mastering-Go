package main

import (
	"container/ring"
	"fmt"
)

var size int = 10

func main() {

	s:=0
	myRing := ring.New(size + 1)
	fmt.Println("Empty ring:", *myRing)

	//  0 - 9
	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		s = s + i
		myRing = myRing.Next()
	}

	fmt.Println(s)
	myRing.Value = 2

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})
	fmt.Println("Sum:", sum)

	for i := 0; i < myRing.Len()+2; i++ {
		myRing = myRing.Next()
		fmt.Print(myRing.Value, " ")
	}
	fmt.Println()
}
