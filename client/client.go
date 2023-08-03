package main

import (
	"context"
	"log"
	"time"

	// replace this line
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := testgrpcapp.NewEchoServiceClient(conn) // replace this line

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &testgrpcapp.EchoMessage{Value: "Hello"}) // replace this line
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Greeting: %s", r.Value)
}
