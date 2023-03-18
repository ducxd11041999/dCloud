package main

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "my-project/dCloud/protobuf/reply"
)

const (
	port = ":7800" // port number to listen on
)

type server struct{}

// ReplyService implement gRPC server
type ReplyService struct{}

// SayHello implement ReplyService
func (s *ReplyService) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in.Message)
	return &pb.Response{Message: "service reply"}, nil
}

func main() {
	// Start gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReplyServer(s, &ReplyService{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
