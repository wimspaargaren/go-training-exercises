package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciClosure(t *testing.T) {
	sequence := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181}
	f := fib()
	for i := 0; i < len(sequence); i++ {
		assert.Equal(t, sequence[i], f())
	}
}
