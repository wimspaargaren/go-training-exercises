package main

import (
	"fmt"
)

var ErrNot42 = fmt.Errorf("it's always 42")

func main() {
	err := ValidateNumber(44)
	_ = err
	// Create a for loop that unwraps the error until there's no underlying error anymore.
}

func ValidateNumber(x int) error {
	if x != 42 {
		return fmt.Errorf("wrap tastic: %w", fmt.Errorf("start to panic: %w", ErrNot42))
	}
	return nil
}
