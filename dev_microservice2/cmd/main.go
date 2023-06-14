// File: dev_microservice2/cmd/main.go

package main

import (
	"log"
	"net"

	"github.com/Nazerkh09/gajap/dev_microservice2/internal/services"
	pb "github.com/Nazerkh09/gajap/dev_microservice2/protobuf"
	"google.golang.org/grpc"
)

func main() {
	// Define the gRPC server
	server := grpc.NewServer()

	// Create an instance of the DocumentService
	documentService := &services.DocumentService{}

	// Register the DocumentService with the server
	pb.RegisterDocumentServiceServer(server, documentService)

	// Start listening on a TCP port
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	log.Printf("gRPC server is listening on port %s", port)

	// Serve incoming gRPC requests
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
