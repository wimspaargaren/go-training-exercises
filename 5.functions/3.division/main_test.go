package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, 11.5, divide(-23, -2))
	assert.Equal(t, -21.0, divide(-42, 2))
	assert.Equal(t, 0.3333333333333333, divide(1, 3))
	assert.True(t, math.IsNaN(divide(42, 0)))
}
