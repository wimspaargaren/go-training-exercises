package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserIDs(t *testing.T) {
	users := []User{{1, "Andy"}, {2, "John"}, {3, "Jane"}}
	userIDs := getUserIDs(users)
	for i := 1; i <= 3; i++ {
		assert.Equal(t, i, *userIDs[i-1])
	}
}
