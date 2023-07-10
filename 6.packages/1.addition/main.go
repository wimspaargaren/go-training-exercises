package main

import (
	"fmt"

	"github.com/wimspaargaren/go-training-exercises/2.packages/1.addition/calculator"
)

func main() {
	// Re-create the Add function, but now move it to a separate package as defined by the import statement above.
	fmt.Println(calculator.Add(21, 21))
}
