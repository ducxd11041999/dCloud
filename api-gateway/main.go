package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	pb "my-project/dCloud/protobuf/reply"
)

func main() {
	// Create gRPC connection
	conn, err := grpc.Dial("localhost:7800", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create gRPC client
	c := pb.NewReplyClient(conn)

	// Create HTTP server with Gorilla Mux
	r := mux.NewRouter()
	r.HandleFunc("/reply", func(w http.ResponseWriter, r *http.Request) {
		// Call gRPC server
		message := &pb.Request{Message: "have request"}
		resp, err := c.SayHello(r.Context(), message)
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Reply: %s", resp.Message)

		// Return response to client
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp.Message))
	})

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":8080", r))
}
