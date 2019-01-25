package server

import (
	"fmt"
	"context"
	"log"
	"net"

	pb "github.com/emakrashov/orca/apiv1pb"
	"github.com/emakrashov/orca/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)


type server struct {
	storage *storage.Storage
}

func (s *server) Set(ctx context.Context, sr *pb.SetRequest) (*pb.SetReply, error) {
	v := []byte(sr.Value)
	s.storage.SetValue(sr.Key, v)
	return &pb.SetReply{Success: true}, nil
}

func (s *server) Get(ctx context.Context, sr *pb.GetRequest) (*pb.GetReply, error) {
	fmt.Println("GET: ", sr.Key)
	v, ok := s.storage.GetValue(sr.Key)
	return &pb.GetReply{Success: ok, Value: string(v)}, nil
}

// LaunchGrpc implements the server launcher.
func LaunchGrpc(storage *storage.Storage) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterApiv1PbServer(s, &server{storage: storage})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
