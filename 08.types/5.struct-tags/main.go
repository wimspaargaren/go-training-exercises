package main

import "encoding/json"

// JSONResponse is a struct with a tagged field for JSON serialisation.
type JSONResponse struct {
	//nolint:govet
	message string `json:"message_as_json_field"`
}

func main() {
	// The result message is empty. Change the JSONResponse struct to fix the issue.
	//nolint:staticcheck
	b, err := json.Marshal(JSONResponse{message: "Hello, World!"})
	if err != nil {
		panic(err)
	}
	// Bonus: See if you can change the program to output yaml, see the following link
	// for common struct tags: https://go.dev/wiki/Well-known-struct-tags
	println(string(b))
}
