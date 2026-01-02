package main

import "fmt"

// Run make lint and see if you can fix the lint errors.
// Hint try running golangci-lint fmt.
func main() {

	fmt.Println(IsTrue(true))
	fmt.Println(IsTrue(false))

}

func IsTrue(x bool) bool {
	if x == true {
		return true
	}
	return false
}
