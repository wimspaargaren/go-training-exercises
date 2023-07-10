package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test42(t *testing.T) {
	tests := []struct {
		Name  string
		Input *X
	}{
		{
			Name: "it's not 99",
			Input: &X{
				V: 99,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			itsAlways42(test.Input)
			assert.Equal(t, 42, test.Input.V)
		})
	}
}
