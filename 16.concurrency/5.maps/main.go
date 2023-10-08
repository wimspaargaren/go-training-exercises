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
	// Hint: use mutual exclusion (sync.Mutex). Also check out what happens if you run the code as is.
	m := NewNonConcurrentMap()
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Add(i, 42)
		}(i)
	}

	wg.Wait()

	fmt.Println("map", m)
}

// NewNonConcurrentMap returns a new ConcurrentMap which is not safe for concurrent use.
func NewNonConcurrentMap() ConcurrentMap {
	return &nonConcurrentMap{m: map[int]int{}}
}

type nonConcurrentMap struct {
	m map[int]int
}

func (m *nonConcurrentMap) Add(k, v int) {
	m.m[k] = v
}
