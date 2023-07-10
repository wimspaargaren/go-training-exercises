package main

import (
	"fmt"
)

var ErrNot42 = fmt.Errorf("it's always 42")

func main() {
	errorList := errorList()
	_ = errorList
	// Create a custom error type that combines the slice of errors
	// to a single error.
	// Once you're done, checkout the newly errors.Join function added in Go 1.20!
}

func errorList() []error {
	errors := []error{}
	for i := 0; i < 10; i++ {
		errors = append(errors, fmt.Errorf("e: %d", i))
	}
	return errors
}
