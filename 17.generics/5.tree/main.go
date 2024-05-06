package main

import "fmt"

// Tree is a generic tree.
type Tree[T any] interface {
	// AddChild adds a child to the tree (i.e. the current node).
	AddChild(Tree[T])
	// Children returns the children of the tree (i.e. the current node).
	Children() []Tree[T]
	// Value returns the value of the tree (i.e. the current node).
	Value() T
}

// NewTree creates a new tree with the given value.
func NewTree[T any](value T) Tree[T] {
	// create the new tree with the given value.
	return nil
}

// implement a generic tree.

func prettyPrintTree[T any](tree Tree[T], indent string) {
	fmt.Printf("%s%v\n", indent, tree.Value())
	for _, child := range tree.Children() {
		prettyPrintTree(child, indent+"---")
	}
}

func main() {
	root := NewTree(42)
	child1 := NewTree(43)
	child2 := NewTree(44)
	root.AddChild(child1)
	root.AddChild(child2)
	child3 := NewTree(45)
	child4 := NewTree(46)
	child1.AddChild(child3)
	child1.AddChild(child4)
	child5 := NewTree(47)
	child6 := NewTree(48)
	child3.AddChild(child5)
	child3.AddChild(child6)
	prettyPrintTree(root, "")
}
