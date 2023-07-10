package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrSlice(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []*int
		Result []int
	}{
		{
			Name: "it's not 99",
			Input: func() []*int {
				x := 41
				x2 := 41
				x3 := 41
				return []*int{&x, &x2, &x3}
			}(),
			Result: []int{42, 42, 42},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			incrSlice(test.Input)
			for i, e := range test.Input {
				assert.Equal(t, test.Result[i], *e)
			}
		})
	}
}
