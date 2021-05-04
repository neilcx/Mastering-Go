package main

import (
	"fmt"
)

type first struct{}

func (a first) F() {
	a.shared()
}

func (a first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	first{}.F() //This is shared() from first!
	second{}.shared() //This is shared() from second!
	i := second{}
	j := i.first
	j.F() //This is shared() from first!
	(i.first).F() //This is shared() from first!
	// (second{}.first).F()
}
