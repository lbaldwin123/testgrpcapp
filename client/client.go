package main

import (
	"context"
	"log"
	"os"
	"time"
	"github.com/lbaldwin123/test-grpc-app"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &echo.EchoMessage{Value: "Hello"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Greeting: %s", r.Value)
}

