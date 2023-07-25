package main

import "fmt"

// Implement a generic (single) linked list of type T.
type Node[T any] interface {
	SetNext(Node[T]) Node[T]
	Next() Node[T]
	Value() T
}

func main() {
	root := NewNode(1)

	next := root.SetNext(NewNode(2))
	for i := 0; i < 10; i++ {
		next = next.SetNext(NewNode(i))
	}

	for n := root; n.Next() != nil; n = n.Next() {
		fmt.Println("Val", n.Value())
	}
}

func NewNode(v int) Node[int] {
	return nil
}
