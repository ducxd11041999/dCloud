package main

import (
	"context"
	"log"
	"net"

	pb "dCloud/protobuf/reply" // import your protobuf package

	"google.golang.org/grpc"
)

const (
	port = ":7800" // port number to listen on
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received request with message: %s", in.Message)
	return &pb.Response{Message: "Hello from A"}, nil
}

func main() {
	// create a gRPC server and register the service
	s := grpc.NewServer()
	pb.RegisterReplyServer(s, &server{})

	// create a listener and start the server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on port %s...", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}