package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Refactor the following code, such that we can call the getTheDocs function
	// without making an actual http request. Hint: can we define an interface for
	// the behaviour of the httpClient and pass it to the function instead of the http client struct?
	// Write a unit test for the getTheDocs function.
	httpClient := http.DefaultClient
	goDocsRespBody := getTheDocs(httpClient)
	fmt.Println(goDocsRespBody)
}

func getTheDocs(client *http.Client) string {
	req, err := http.NewRequest(http.MethodGet, "https://go.dev/doc/", nil)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	resp, err := client.Do(req)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		// we did not learn about errors yet, so ignore for now
		return err.Error()
	}
	return string(b)
}
