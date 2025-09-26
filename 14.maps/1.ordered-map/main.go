package main

import "fmt"

// Implement the following map interface.
type OrderedMap interface {
	// Add adds a key value pair to the map
	Add(key string, value string)
	// GetValuesOrdered returns the values of the ordered map in the same order as they were inserted.
	// If a value with the same key was added twice, the ordering remains the same.
	// Hint: Use a slice to store the keys in the order they were added.
	GetValuesOrdered() []string
}

func GetOrderedMap() OrderedMap {
	// FIXME: Implement; Run the tests to verify your solution.
	return nil
}

func main() {
	orderedMap := GetOrderedMap()
	orderedMap.Add("k1", "v1")
	orderedMap.Add("k2", "v2")
	orderedMap.Add("k3", "v3")
	orderedMap.Add("k1", "v1")
	orderedMap.Add("k3", "v3")

	fmt.Println(orderedMap.GetValuesOrdered())
}
