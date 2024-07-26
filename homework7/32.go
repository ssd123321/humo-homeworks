package main

import "fmt"

func ThreeNodes() {
	var n1, n2, n3 Node
	n1.Value = 1
	n2.Value = 2
	n3.Value = 3
	n1.Next = &n2
	n2.Next = &n3
	current := &n1
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}

}
