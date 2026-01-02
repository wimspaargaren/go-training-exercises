package main

import "fmt"

// implement the fibonacci using a closure.
func fib() func() int {
	return nil
}

func main() {
	f := fib()
	for range 10 {
		fmt.Println(f())
	}
}
