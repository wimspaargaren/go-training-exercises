package main

import "fmt"

// Stack Implement a generic stack where we can push and pop elements of type T.
type Stack[T any] interface {
	Push(v T)
	Pop() T
}

func main() {
	intStack := NewIntStack()
	intStack.Push(42)
	fmt.Println(intStack.Pop())

	stringStack := NewStringStack()
	stringStack.Push("42")
	fmt.Println(stringStack.Pop())
}

func NewIntStack() Stack[int] {
	return nil
}

func NewStringStack() Stack[string] {
	return nil
}
