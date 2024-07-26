package main

func AddNode(n1 *Node) {
	var numberx Node = Node{37, nil}
	current := n1
	for {
		if current.Next != nil {
			current = current.Next
			continue
		} else {
			current.Next = &numberx
			break
		}
	}
}
