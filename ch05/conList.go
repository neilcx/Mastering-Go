package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}

	fmt.Println()
}

func main() {

	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	fmt.Printf("====\n")
	printList(values)

	fmt.Printf("====\n")
	values.PushFront("Three")

	fmt.Printf("====\n")
	printList(values)

	fmt.Printf("====\n")
	values.InsertBefore("Four", e1)
	fmt.Printf("====\n")
	printList(values)

	fmt.Printf("====\n")
	values.InsertAfter("Five", e2)
	fmt.Printf("====\n")
	printList(values)

	fmt.Printf("====\n")
	values.Remove(e2)
	values.Remove(e2)
	//printList(values)
	values.InsertAfter("FiveFive", e2)
	printList(values)
	values.PushBackList(values)

	printList(values)

	values.Init()
	fmt.Printf("After Init(): %v\n", values)

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
}

