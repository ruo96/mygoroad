package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server instance
	server := grpc.NewServer()

	// ... register server handlers ...

	// Start the server on port 8080
	port := ":18080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
