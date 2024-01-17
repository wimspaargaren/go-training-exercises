package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCalculationFunction(t *testing.T) {
	tests := []struct {
		Name   string
		Input  CalcType
		Output int
	}{
		{
			Name:   "Calc type sum",
			Input:  CalcTypeSum,
			Output: 42,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			calcFunc := GetCalculationFunction(test.Input)
			assert.Equal(t, test.Output, calcFunc(21, 21))
		})
	}
}
