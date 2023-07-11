package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	tests := []struct {
		Name   string
		Input  string
		Output map[string]int
	}{
		{
			Name:  "empty array",
			Input: "foo foo bar foo bar bar foo",
			Output: map[string]int{
				"foo": 4,
				"bar": 3,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			wordCountMap := CountWords(test.Input)

			assert.Equal(t, test.Output, wordCountMap)

		})
	}
}
