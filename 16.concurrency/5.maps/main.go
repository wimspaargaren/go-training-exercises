package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap interface {
	Add(k, v int)
}

func main() {
	// Maps do not support concurrent writes. Implement the ConcurrentMap type
	// such that we can concurrently assign values to our map.
	// Hint: use mutual exclusion.
	m := map[int]int{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = 42
		}(i)
	}

	wg.Wait()

	fmt.Println("map", m)
}
