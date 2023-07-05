package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 42, add(21, 21))
	assert.Equal(t, -42, add(-21, -21))
	assert.Equal(t, 0, add(0, 0))
}
