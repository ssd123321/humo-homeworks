package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func TwoNodes() {
	var n1, n2 Node
	n1.Value = 5
	n2.Value = 8
	n1.Next = &n2
	repeat := &n1
	for repeat != nil {
		fmt.Println(repeat.Value)
		repeat = repeat.Next
	}
}
