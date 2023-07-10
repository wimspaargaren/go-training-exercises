package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	tests := []struct {
		Name       string
		InputArray []int
		InputElem  int
		Output     bool
	}{
		{
			Name: "empty array",
		},
		{
			Name:       "does not exist",
			InputArray: []int{1},
			InputElem:  42,
		},
		{
			Name:       "does  exist",
			InputArray: []int{42},
			InputElem:  42,
			Output:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Exists(test.InputArray, test.InputElem)
			assert.Equal(t, test.Output, out)
		})
	}
}
