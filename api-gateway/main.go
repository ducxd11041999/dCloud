package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	pb "dCloud/protobuf/reply" // import your protobuf package
)

const (
	port      = ":8081" // port number to listen on
	grpcPort  = ":7800" // port number of gRPC server
	endpointA = "localhost" + grpcPort
	endpointB = "localhost" + grpcPort
)


func main() {
	// create a context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// create a new gRPC gateway mux and register endpoints
	gwmux := runtime.NewServeMux()

	// register endpoint helloA
	err := pb.NewReplyClient(ctx, gwmux, endpointA, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("failed to register endpoint helloA: %v", err)
	}

	// register endpoint helloB
	err = pb.(ctx, gwmux, endpointB, []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("failed to register endpoint helloB: %v", err)
	}

	// create a new HTTP server and listen on port 8081
	log.Printf("starting HTTP server on port %s", port)
	if err := http.ListenAndServe(port, gwmux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}