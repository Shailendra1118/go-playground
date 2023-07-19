package main

import (
	"context"
	"fmt"
	"go-playground/greet/greetpb/go_playground/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	//"log"
)

type server struct {
	proto.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *proto.GreetingRequest) (*proto.GreetingResponse, error) {
	fmt.Printf("Greet function is invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	return &proto.GreetingResponse{
		Result: result,
	}, nil

}

func main() {
	fmt.Println("Hello Shailendra!")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
