package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookupNode(t *Node, v int) bool {
	if root == nil {

		// ?
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root) //&{0 <nil>}
	root = nil
	traverse(root) //-> Empty list!"
	addNode(root, 1)
	addNode(root, -1)
	traverse(root) //1 -> -1 ->
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5) //Node already exists: 5
	addNode(root, 5) //Node already exists: 5
	traverse(root) //1 -> -1 -> 10 -> 5 -> 45 ->
	addNode(root, 100)
	traverse(root) //1 -> -1 -> 10 -> 5 -> 45 ->100 ->


	root = nil
	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, 100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}
}
