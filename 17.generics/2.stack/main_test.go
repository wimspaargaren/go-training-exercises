package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenericStack(t *testing.T) {
	intStack := NewIntStack()
	intStack.Push(10)
	intStack.Push(20)
	assert.Equal(t, 20, intStack.Pop())
	assert.Equal(t, 10, intStack.Pop())
	assert.Equal(t, 0, intStack.Pop())
}
