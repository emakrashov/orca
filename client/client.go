package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/emakrashov/orca/apiv1pb"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func clientGet(ctx context.Context, c pb.Apiv1PbClient) {
	if len(os.Args) <= 2 {
		fmt.Println("Usage:")
		fmt.Println("orca-client get a")
		os.Exit(1)
	}
	v, err := c.Get(ctx, &pb.GetRequest{Key: os.Args[2]})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Fetched: %s", v.Value)
}

func clientSet(ctx context.Context, c pb.Apiv1PbClient) {
	if len(os.Args) <= 3 {
		fmt.Println("Usage:")
		fmt.Println("orca-client set a 5")
		os.Exit(1)
	}
	value := os.Args[3]
	_, err := c.Set(ctx, &pb.SetRequest{Key: os.Args[2], Value: value})
	if err != nil {
		log.Fatalf("could not set", err)
	}
	log.Printf("Success")
}

// orca-client set a 5
// orca-client get a
func main() {
	if len(os.Args) <= 1 {
		// name = os.Args[1]
		fmt.Println("Usage:")
		fmt.Println("orca-client set a 5")
		fmt.Println("orca-client get a")
		os.Exit(0)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	cli := pb.NewApiv1PbClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch os.Args[1] {
	case "get":
		clientGet(ctx, cli)
	case "set":
		clientSet(ctx, cli)
	}
}
