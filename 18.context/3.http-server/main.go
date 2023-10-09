package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(authorizationMiddleware)
	router.HandleFunc("/user", handleUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":9000", router))
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
