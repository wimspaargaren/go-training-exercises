package main

import "fmt"

// Implement the following map interface
type OrderderMap interface {
	// Add adds a key value pair to the map
	Add(key string, value string)
	// GetValuesOrdered returns the values of the ordered map in the same order as they were inserted.
	// If a value with the same key was added twice, the ordering remains the same.
	GetValuesOrdered() []string
}

func GetOrderdMap() OrderderMap {
	// FIXME: Implement; Run the tests to verify your solution.
	return nil
}

func main() {
	orderedMap := GetOrderdMap()
	orderedMap.Add("k1", "v1")
	orderedMap.Add("k2", "v2")
	orderedMap.Add("k3", "v3")
	orderedMap.Add("k1", "v1")
	orderedMap.Add("k3", "v3")

	fmt.Println(orderedMap.GetValuesOrdered())
}
