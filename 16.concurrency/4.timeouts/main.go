package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	result, err := timeIsMoney()
	if err != nil {
		log.Fatalf("function timed out: %s", err)
	}
	fmt.Println("Result:", result)
}

// Implement the timeIsMoney function. The  function performs some asynchronous
// processing. Once it's done it will put the result on the channel. If the processing
// takes longer than 3 seconds, the function should return an error. Otherwise, the function
// should return the value
func timeIsMoney() (int, error) {
	ch := make(chan int, 1)

	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(r.Intn(5000) * int(time.Millisecond)))
		ch <- 42
	}()
	// Hint: Use a select statement to listen to multiple channels
	// Hint: Use the time.After channel to timeout after a given duration

	return 0, nil
}
