package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Define multiple test cases using table tests.
	tests := []struct {
		Name   string
		X, Y   int
		Output int
	}{}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Add(test.X, test.Y)
			assert.Equal(t, test.Output, out)
		})
	}
}
