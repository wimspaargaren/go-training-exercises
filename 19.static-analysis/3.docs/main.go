package main

import (
	"fmt"

	"github.com/wimspaargaren/go-training-exercises/19.static-analysis/3.docs/mypkg"
)

// Run make lint and see if you can fix the lint errors.
func main() {
	fmt.Println(mypkg.HelloWorld())
}
