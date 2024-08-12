package main

import "fmt"

type PointerOperations interface {
	Increment()
	Decrement()
}

type IntPointer struct {
	value *int
}

func (p *IntPointer) Increment() {
	*p.value++
}

func (p *IntPointer) Decrement() {
	*p.value--
}

func performOperations(p PointerOperations) {
	p.Increment()
	p.Decrement()
	fmt.Println("Value:", *p.(*IntPointer).value)
}
