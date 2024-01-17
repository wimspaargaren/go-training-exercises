package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPointer(t *testing.T) {
	tests := []struct {
		Name   string
		InputX *int
		InputY int
		Output int
	}{
		{
			Name: "add 1",
			InputX: func() *int {
				x := 42
				return &x
			}(),
			InputY: 1,
			Output: 43,
		},
		{
			Name: "add negative",
			InputX: func() *int {
				x := 42
				return &x
			}(),
			InputY: -42,
			Output: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			addPointer(test.InputX, test.InputY)
			assert.Equal(t, test.Output, *test.InputX)
		})
	}
}
