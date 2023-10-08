package main

func main() {
	go func(a, b int) {
		c := a + b
		_ = c
	}(1, 2)
	// Retrieve the computed value c from the goroutine
	// and print it in the main function.
	// Hint: use a channel to retrieve the value.
}
