package main

import (
	"context"
	"log"

	pb "github.com/YuryZhehala/grpc-go-cource/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Clement"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
