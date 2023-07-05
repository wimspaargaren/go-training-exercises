package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Finish the following program, such that it prints the smallest number.
	x := s1.Intn(100)
	y := s1.Intn(100)
	fmt.Println(x, y)
}
