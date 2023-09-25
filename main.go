package main

import (
	"fmt"
)

func main() {
	sayHello("Hello Gophers!")
}

func sayHello(someText string) {
	fmt.Println(someText)
}
