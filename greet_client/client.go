package main

import (
	"context"
	"fmt"
	"go-playground/greet/greetpb/go_playground/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I am Client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := proto.NewGreetServiceClient(conn)
	fmt.Println("greeling client is up", c)

	doUnary(c)

}

func doUnary(c proto.GreetServiceClient) {
	fmt.Println("Starting to do unary RPC call...")
	res, err := c.Greet(context.Background(), &proto.GreetingRequest{
		Greeting: &proto.Greeting{
			FirstName: "Shailendra",
			LastName:  "Yadav",
		},
	})

	if err != nil {
		log.Fatalf("Error while calling Greet RPC %v", err)
	}

	log.Printf("Response from server: %v", res.Result)
}
