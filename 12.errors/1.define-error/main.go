package main

import (
	"fmt"
)

// Define a error type ErrNegativeSqrt float64
// Implement the Sqrt function such that it return the Sqrt for a
// given floating point number.
// In case the input number is negative, the function should
// return the ErrNegativeSqrt.

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
