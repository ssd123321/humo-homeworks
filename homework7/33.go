package main

import "fmt"

func GetAllValues(n1 *Node) {
	current := n1
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
