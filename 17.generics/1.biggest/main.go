package main

import (
	"cmp"
	"fmt"
)

func main() {
	fmt.Println(Max[uint](3, 4))
	fmt.Println(Max[float64](3, 4))
	fmt.Println(Max[int](3, 4))
}

// Max Implement the generic Max function, which returns the largest input value.
func Max[T cmp.Ordered](x, y T) T {
	return x
}
