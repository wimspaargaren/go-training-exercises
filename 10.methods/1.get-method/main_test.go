package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArea(t *testing.T) {
	tests := []struct {
		Name     string
		Triangle triangle
		Expected float64
	}{
		{
			Name: "Base 12 Height 6",
			Triangle: triangle{
				base:   12,
				height: 6,
			},
			Expected: 36,
		},
		{
			Name: "Base 15 Height 8",
			Triangle: triangle{
				base:   15,
				height: 8,
			},
			Expected: 60,
		},
		{
			Name: "0 values",
			Triangle: triangle{
				base:   0,
				height: 0,
			},
			Expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := test.Triangle.Area()
			assert.Equal(t, test.Expected, actual)
		})
	}
}
