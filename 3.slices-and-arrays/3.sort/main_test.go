package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{
			Name: "empty array",
		},
		{
			Name:   "ordering single letter",
			Input:  []string{"b", "a", "c"},
			Output: []string{"a", "b", "c"},
		},
		{
			Name:   "ordering multi letter",
			Input:  []string{"ab", "aa", "aaa"},
			Output: []string{"aa", "aaa", "ab"},
		},
		{
			Name:   "ordering unicode",
			Input:  []string{"😀", "b", "a", "c"},
			Output: []string{"a", "b", "c", "😀"},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			out := Sort(test.Input)
			assert.Equal(t, test.Output, out)
		})
	}
}
