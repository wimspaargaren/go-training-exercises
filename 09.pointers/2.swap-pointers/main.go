package main

import "fmt"

// 1 & 2. Define the swap function.
//
//nolint:unused
func swap(a *int, b *int) {
	// Implement the swap logic here using the dereference operator (*)
}

func main() {
	// 3. Declare and initialise variables
	num1 := 5
	num2 := 10

	fmt.Println("Before swap: num1 =", num1, "num2 =", num2)

	// 4. Call swap, passing the addresses
	// swap(...)

	fmt.Println("After swap: num1 =", num1, "num2 =", num2)
}
