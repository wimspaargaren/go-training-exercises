package main

// Exercise: Implement a gRPC client for the Greeter service.
//
// 1. Ensure the protobuf code has been generated (run "make generate" in the parent directory).
// 2. Implement the client below:
//    - Connect to the gRPC server at localhost:50051.
//    - Call SayHello with your name and print the response.
//    - Call SayGoodbye with your name and print the response.
// 3. Start the server first, then run this client.

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"
//
// 	pb "github.com/wimspaargaren/go-training-exercises/21.protobuf/4.grpc-service/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

func main() {
	// TODO: Uncomment and complete the client:

	// conn, err := grpc.NewClient("localhost:50051",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalf("failed to connect: %v", err)
	// }
	// defer conn.Close()
	//
	// client := pb.NewGreeterClient(conn)
	//
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	//
	// // TODO: Call SayHello and print the response
	// helloResp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})
	// if err != nil {
	// 	log.Fatalf("SayHello failed: %v", err)
	// }
	// fmt.Println(helloResp.GetMessage())
	//
	// // TODO: Call SayGoodbye and print the response
	// goodbyeResp, err := client.SayGoodbye(ctx, &pb.HelloRequest{Name: "Gopher"})
	// if err != nil {
	// 	log.Fatalf("SayGoodbye failed: %v", err)
	// }
	// fmt.Println(goodbyeResp.GetMessage())
}
