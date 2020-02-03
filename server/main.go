package main

import (
	"context"
	"log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	"github.com/work/rails-grpc/server/helloworld"
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
