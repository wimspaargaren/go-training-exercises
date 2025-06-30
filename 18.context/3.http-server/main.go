package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/user", authorizationMiddleware(http.HandlerFunc(handleUser)))

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

func handleUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("uuid")
	if userID == nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Who are you?"))
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Printf("[%v] Returning 200 - ", userID)
	w.Write([]byte(fmt.Sprintf("You are user: %s", userID)))
}

func authorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := uuid.New()
		_ = userID
		// Add the user ID to the request context
		// Note: Just setting a value in a context, does not add it to the parent context.
		// The *http.Request struct has a WithContext method that returns a copy of the request
		next.ServeHTTP(w, r)
	})
}
