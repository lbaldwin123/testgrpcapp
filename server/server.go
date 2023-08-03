package main

import (
	"context"
	"log"
	"net"

	"github.com/lbaldwin123/testgrpcapp" // replace this line
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *testgrpcapp.EchoMessage) (*testgrpcapp.EchoMessage, error) {
	return &testgrpcapp.EchoMessage{Value: "Echo: " + in.Value}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	testgrpcapp.RegisterEchoServiceServer(s, &server{}) // replace this line
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
