package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Define a context value key type

func main() {
	ctx := context.Background()
	// Use a context with timeout to wait not longer than 5 seconds until the result is returned
	// otherwise log an error message that the operation timed out.
	resChan := make(chan (int))

	go flakyWork(ctx, resChan)

	x := <-resChan

	fmt.Println("X", x)
}

func flakyWork(ctx context.Context, resChan chan (int)) {
	defer close(resChan)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(r.Intn(10) * int(time.Second)))
	resChan <- 42
}
