package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{
			Name: "empty array",
		},
		{
			Name:   "No duplicate entrires",
			Input:  []string{"1", "2", "3", "4"},
			Output: []string{"1", "2", "3", "4"},
		},
		{
			Name:   "Duplicate entries at the start",
			Input:  []string{"1", "2", "3", "4", "2", "3"},
			Output: []string{"1", "2", "3", "4"},
		},
		{
			Name:   "Duplicate entries at the end",
			Input:  []string{"1", "1", "3", "4", "2", "3"},
			Output: []string{"1", "3", "4", "5"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			orderedMap := GetOrderdMap()
			for i, v := range test.Input {
				orderedMap.Add(v, fmt.Sprintf("%d", i+1))
			}
			assert.Equal(t, test.Output, orderedMap.GetValuesOrdered())

		})
	}
}
