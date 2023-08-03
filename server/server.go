package main

import (
	"context"
	"log"
	"net"

	"github.com/lbaldwin123/test-grpc-app"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *echo.EchoMessage) (*echo.EchoMessage, error) {
	return &echo.EchoMessage{Value: "Echo: " + in.Value}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

