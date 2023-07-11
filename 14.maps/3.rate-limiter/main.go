package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Implement the following RateLimiter interface
type RateLimiter interface {
	// ShouldLimit function takes an id and returns if the given id should be
	// rate limited or not. To keep it simple, we'll just return true as soon
	// as a given id has visited more than 5 times.
	ShouldLimit(id uuid.UUID) bool
}

func NewRateLimiter() RateLimiter {
	return nil
}

func main() {
	rateLimiter := NewRateLimiter()

	userID := uuid.New()
	for i := 0; i < 3; i++ {
		rateLimiter.ShouldLimit(userID)
	}

	fmt.Println("is user rate limited", rateLimiter.ShouldLimit(userID))
	fmt.Println("is user rate limited", rateLimiter.ShouldLimit(userID))
}
