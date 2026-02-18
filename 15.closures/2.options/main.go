package main

import (
	"fmt"
	"time"
)

type options struct {
	requestTimeout time.Duration
	sslEnforced    bool
}

func defaultOptions() *options {
	return &options{
		requestTimeout: time.Second * 10,
		sslEnforced:    true,
	}
}

type httpClient struct {
	options *options
}

type OptFunc func(*options)

func NewHTTPClient(opts ...OptFunc) httpClient {
	defaultOpts := defaultOptions()
	for _, f := range opts {
		f(defaultOpts)
	}
	return httpClient{
		options: defaultOpts,
	}
}

func main() {
	// Define closures returning the OptFunc to adjust the requestTimeout
	// and enforceSSL options. Afterwards call the NewHTTPClient with your
	// OptFuncs.
	fmt.Println(NewHTTPClient().options)
}
