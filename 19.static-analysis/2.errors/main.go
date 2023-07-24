package main

import "fmt"

// Run make lint and see if you can fix the lint errors.
func main() {
	WhatCouldGoWrong()
}

func WhatCouldGoWrong() error {
	return fmt.Errorf("very error")
}
