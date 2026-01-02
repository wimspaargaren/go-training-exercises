package main

import (
	"fmt"
)

// Cache Implement the generic cache interface.
type Cache[T comparable, U any] interface {
	Set(k T, v U)
	Get(k T) U
}

func main() {
	cache := NewIntCache()
	cache.Set(42, "hello cache")
	fmt.Println(cache.Get(42))
}

func NewIntCache() Cache[int, string] {
	return nil
}
