package main

import pb "github.com/YuryZhehala/grpc-go-cource/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
