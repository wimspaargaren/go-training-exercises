package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMax(t *testing.T) {
	tests := []struct {
		Name      string
		Input     []int
		OutputMin int
		OutputMax int
	}{
		{
			Name: "empty array",
		},
		{
			Name:      "one element",
			Input:     []int{42},
			OutputMin: 42,
			OutputMax: 42,
		},
		{
			Name:      "multiple elements with negative number",
			Input:     []int{42, 0, -2},
			OutputMin: -2,
			OutputMax: 42,
		},
		{
			Name:      "multiple elements without negative number",
			Input:     []int{3, 2, 1},
			OutputMin: 1,
			OutputMax: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			minOut, maxOut := MinMax(test.Input)
			assert.Equal(t, test.OutputMin, minOut)
			assert.Equal(t, test.OutputMax, maxOut)
		})
	}
}
