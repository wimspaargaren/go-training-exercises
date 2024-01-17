package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Repeat the previous assignment using a switch statement.
	x := s1.Intn(100)
	y := s1.Intn(100)
	fmt.Println(x, y)
}
