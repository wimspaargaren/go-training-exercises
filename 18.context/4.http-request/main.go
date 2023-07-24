package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Consider the following HTTP server; Write a Go program which performs an http request to the server,
// but uses a context which ensures we use a maximum of 5 seconds to wait for an answer.
// You can either bootstrap the client from the same main.go in a go routine, or create a
// sub folder with a new main.go
// Hint: use http.NewRequestWithContext
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/flaky-endpoint", flakyHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9000", router))
}

func flakyHandler(w http.ResponseWriter, r *http.Request) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Second)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Flaky endpoint!"))
}
