package main

import "fmt"

// 1 & 2. Define the safePrint function.
func safePrint(p *string) {
	// Use an if p != nil check here
	// ...
	fmt.Println(*p)
}

func main() {
	// Test 1: nil pointer
	var p1 *string
	safePrint(p1)

	// Test 2: using new()
	// 3. Use new() to create pointer 2
	// p2 := new(...)

	// 3. Assign a value to the underlying string
	// *p2 = ...

	// safePrint(p2)
}
