package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Nazerkh09/gajap/dev_microservice1/internal"
	pb "github.com/Nazerkh09/gajap/dev_microservice1/protobuf"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

func main() {
	// Create a new RabbitMQ connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a new RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	// Declare the RabbitMQ queue
	queueName := "authentication_queue"
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare RabbitMQ queue: %v", err)
	}

	// Create a new gRPC server
	server := grpc.NewServer()

	// Register the authentication service with the server
	authService := internal.NewAuthenticationService()
	pb.RegisterAuthenticationServiceServer(server, authService)

	// Start the gRPC server in a separate goroutine
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		log.Println("Starting Authentication gRPC server...")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	log.Println("Authentication microservice is running.")

	// Gracefully handle termination signals
	waitForTerminationSignal()
}

func waitForTerminationSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Println("Shutting down...")
	server.GracefulStop()
}
