package main

import (
	"fmt"
	"time"
)

type options struct {
	requestTimout time.Duration
	sslEnforced   bool
}

func defaultOptions() *options {
	return &options{
		requestTimout: time.Second * 10,
		sslEnforced:   true,
	}
}

type httpClient struct {
	options *options
}

type OptFunc func(*options)

func WithEnableSSLOption() OptFunc {
	return func(o *options) {
		o.sslEnforced = true
	}
}

func WithDurationOption(d time.Duration) OptFunc {
	return func(o *options) {
		o.requestTimout = d
	}
}

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
	NewHTTPClient(
		WithEnableSSLOption(),
		WithDurationOption(time.Second*5),
	)

	NewHTTPClient(func(*options) {})
	NewHTTPClient(func(*options) {}, func(o *options) {})
	optsFuncList := []OptFunc{
		func(o *options) {
		},
	}
	NewHTTPClient(optsFuncList...)

}
