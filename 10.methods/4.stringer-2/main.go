package main

import "fmt"

// Define USD as a new type based on float64
type USD float64

// Implement the String() method for the USD type.
// It should have a value receiver `(m USD)`.
// It should return a string.
// Hint: Use fmt.Sprintf() to format the string.
func (m USD) String() string {
	// Your code here
	return ""
}

func main() {
	salary := USD(50000.75)
	bonus := USD(10000.25)

	fmt.Println("My salary is:", salary)
	fmt.Println("My bonus is:", bonus)
	fmt.Println("Total compensation:", salary+bonus)
}
