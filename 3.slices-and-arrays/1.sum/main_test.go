package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		Name   string
		Input  [4]int
		Output int
	}{
		{
			Name: "empty array",
		},
		{
			Name:   "negative numbers",
			Input:  [4]int{-1, -2, -3, -4},
			Output: -10,
		},
		{
			Name:   "negative numbers",
			Input:  [4]int{1, 2, 3, 4},
			Output: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Sum(test.Input)
			assert.Equal(t, test.Output, out)
		})
	}
}
