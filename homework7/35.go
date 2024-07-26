package main

func DeleteNode(n1 *Node, Value int) {
	if n1.Value == Value {
		*n1 = *n1.Next
	} else if n1.Next.Value == Value {
		n1.Next = n1.Next.Next
	} else {
		current := n1
		current1 := n1.Next
		for current1 != nil {
			if current1.Value == Value {
				current.Next = current1.Next
			}
			current = current.Next
			current1 = current1.Next
		}
	}

}
