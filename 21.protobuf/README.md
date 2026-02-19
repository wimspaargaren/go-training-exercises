# 21. Protocol Buffers (Protobuf) in Go

These exercises cover working with Protocol Buffers and gRPC in Go.

## Prerequisites

Install the `protoc` compiler and Go plugins:

```bash
# Install protoc (macOS)
brew install protobuf

# Install Go plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$GOPATH/bin` is in your `PATH`:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Exercises

Each exercise includes a `Makefile` to generate Go code from `.proto` files.
Run `make generate` in the exercise directory before working on the Go code.

1. **basic-message** - Define a protobuf message and use it in Go.
2. **marshal-unmarshal** - Serialize and deserialize protobuf messages.
3. **nested-messages** - Work with nested messages, repeated fields, and enums.
4. **grpc-service** - Implement a gRPC server and client.
