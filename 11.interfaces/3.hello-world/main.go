package main

type Stringer interface {
	String() string
}

func main() {
	var s Stringer
	// Implement the stringer interface and pass it to the log function.
	log(s)
}

func log(s Stringer) {
	println(s.String())
}
