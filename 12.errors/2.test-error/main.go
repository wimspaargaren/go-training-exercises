package main

import (
	"fmt"
)

// Sqrt Copy your solution from the previous assignment and write a unit tests covering all branches and matching the
// specific ErrNegativeSqrt result. Hint: you can leverage the assert.ErrorAs function from testify.
func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
