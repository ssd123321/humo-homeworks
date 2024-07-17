package main

func Add(n1, n2 int) int {
	if n1 < 0 {
		n1 *= -1
	}
	if n2 < 0 {
		n2 *= -1
	}
	return n1 + n2
}
