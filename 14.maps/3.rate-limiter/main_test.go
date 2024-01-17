package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter_ShouldLimit(t *testing.T) {
	rl := NewRateLimiter()

	// Test case 1: Limit not reached
	id1 := uuid.New()
	for i := 0; i < 5; i++ {
		assert.False(t, rl.ShouldLimit(id1))
	}

	// Test case 2: Limit reached
	assert.True(t, rl.ShouldLimit(id1))

	// Test case 3: Limit not reached for different IDs
	id2 := uuid.New()
	for i := 0; i < 5; i++ {
		assert.False(t, rl.ShouldLimit(id2))
	}

	// Test case 4: Limit reached for different IDs
	assert.True(t, rl.ShouldLimit(id2))
}
