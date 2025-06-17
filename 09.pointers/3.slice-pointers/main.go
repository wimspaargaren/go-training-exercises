package main

import "fmt"

// This function takes a slice (which is a reference type).
func modifySlice(s []int) {
	if len(s) > 0 {
		// Your code here: change the first element to 999
	}
}

// This function must take a pointer to a slice to modify
// the slice header (length/capacity) of the original slice variable.
func growSlice(s *[]int) {
	// Your code here: append the number 100 to the slice
}

func main() {
	mySlice := []int{1, 2, 3}
	fmt.Printf("Original slice: %v, len: %d, cap: %d\n", mySlice, len(mySlice), cap(mySlice))

	modifySlice(mySlice)
	fmt.Printf("After modifySlice: %v, len: %d, cap: %d\n", mySlice, len(mySlice), cap(mySlice))

	growSlice(&mySlice)
	fmt.Printf("After growSlice: %v, len: %d, cap: %d\n", mySlice, len(mySlice), cap(mySlice))
}
