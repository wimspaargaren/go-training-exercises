package main

import "fmt"

// implement the fibonacci using a closure.
func fib() func() int {
	return nil
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
