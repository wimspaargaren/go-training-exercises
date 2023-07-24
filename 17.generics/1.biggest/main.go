package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Max[uint](3, 4))
	fmt.Println(Max[float64](3, 4))
	fmt.Println(Max[int](3, 4))
}

// Implement the generic Max function, which returns the largest input value.
func Max[T constraints.Ordered](x, y T) T {
	return x
}
