package main

// Exercise: Implement a gRPC server for the Greeter service.
//
// 1. Run "make generate" in the parent directory to generate the gRPC Go code.
// 2. Implement the GreeterServer below:
//    - SayHello should return "Hello, <name>!"
//    - SayGoodbye should return "Goodbye, <name>!"
// 3. Start the gRPC server on port :50051.
// 4. Test it by running the client in the client/ directory.

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net"
//
// 	pb "github.com/wimspaargaren/go-training-exercises/21.protobuf/4.grpc-service/pb"
// 	"google.golang.org/grpc"
// )

// type server struct {
// 	pb.UnimplementedGreeterServer
// }

// TODO: Implement SayHello - should return "Hello, <name>!"
// func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	return nil, nil
// }

// TODO: Implement SayGoodbye - should return "Goodbye, <name>!"
// func (s *server) SayGoodbye(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
// 	return nil, nil
// }

func main() {
	// TODO: Uncomment and complete the server setup:

	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	//
	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	//
	// fmt.Println("gRPC server listening on :50051")
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
