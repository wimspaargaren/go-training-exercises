package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Consider the following HTTP server; Write a Go program which performs an http request to the server,
// but uses a context which ensures we use a maximum of 5 seconds to wait for an answer.
// You can either bootstrap the client from the same main.go in a go routine, or create a
// sub folder with a new main.go
// Hint: use http.NewRequestWithContext
func main() {
	mux := http.NewServeMux()
	mux.Handle("/flaky-endpoint", http.HandlerFunc(flakyHandler))

	server := &http.Server{
		Handler: mux,
		Addr:    ":9000",
	}
	go func() {
		log.Println("Starting server on :9000")
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed")
		} else if err != nil {
			log.Fatalf("Could not start server: %v", err)
		}
	}()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server gracefully stopped")
}

func flakyHandler(w http.ResponseWriter, r *http.Request) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Second)))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Flaky endpoint!"))
}
