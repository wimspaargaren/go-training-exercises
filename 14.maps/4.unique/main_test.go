package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output int
	}{
		{
			Name: "empty array",
		},
		{
			Name:   "slice with duplicates",
			Input:  []string{"1", "2", "2", "3", "4", "4"},
			Output: 4,
		},
		{
			Name:   "slice without duplicates",
			Input:  []string{"1", "2", "3", "4"},
			Output: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			wordCountMap := UniqueElements(test.Input)
			assert.Equal(t, test.Output, wordCountMap)
		})
	}
}
