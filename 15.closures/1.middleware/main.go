package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Try starting the program and execute an HTTP request using curl: curl localhost:3000/home
func main() {
	fmt.Println("Start HTTP server")

	// server := http.NewServeMux()
	// server.

	http.HandleFunc("/home", userAgentLogger("foo", middlewareNew(homepageHandler)))

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Define a function returning a closure, which prints the user agent of the request.
// Hint: both the input parameter and return type of the function should adhere to the http Handler function which
// looks as follows: func(http.ResponseWriter, *http.Request)

func homepageHandler(x Request) {
	fmt.Println("x", x)
	fmt.Fprintln(x.w, "Hello from server")
	fmt.Fprintln(x.w, "I am learning about closures")
}

type Request struct {
	r *http.Request
	w http.ResponseWriter
}

type MyNewAPIHandler func(Request)

func userAgentLogger(name string, helloWorld MyNewAPIHandler) func(http.ResponseWriter, *http.Request) {
	fmt.Println("hiiii", name)
	return func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Println("User agent:", r.UserAgent())
		helloWorld(Request{r, w})
		fmt.Println("hello world is called", time.Since(now))
	}
}

func middlewareNew(x MyNewAPIHandler) MyNewAPIHandler {
	return func(r Request) {
		fmt.Println("middleware called")
		x(r)
	}
}
