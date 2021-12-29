package main

import (
	"context"
	"log"
	"net"

	"github.com/pobyzaarif/belajarGo/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &proto.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterCalculateServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
