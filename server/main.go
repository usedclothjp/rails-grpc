package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/work/rails-grpc/server/helloworld"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	name := req.GetName()
	fmt.Println(name)

	return &helloworld.HelloReply{
		Message: name,
	}, nil
}

func main() {
	fmt.Println("server start")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
