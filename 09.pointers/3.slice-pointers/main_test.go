package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGrowSlice(t *testing.T) {
	tests := []struct {
		Name   string
		Input  *[]int
		Result []int
	}{
		{
			Name: "it's not 99",
			Input: func() *[]int {
				return &[]int{41, 41, 41}
			}(),
			Result: []int{41, 41, 41, 100},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			growSlice(test.Input)
			require.Equal(t, len(test.Result), len(*test.Input), "length of slice should match")
			for i, e := range test.Result {
				assert.Equal(t, (*test.Input)[i], e)
			}
		})
	}
}

func TestModifySlice(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []int
		Result []int
	}{
		{
			Name: "modify first element",

			Input:  []int{1, 2, 3},
			Result: []int{999, 2, 3},
		},
		{
			Name:   "empty slice",
			Input:  []int{},
			Result: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			modifySlice(test.Input)
			assert.Equal(t, test.Result, test.Input)
		})
	}
}
