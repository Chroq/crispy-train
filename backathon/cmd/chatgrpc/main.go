package main

import (
	"backathon/proto/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Chat struct {
	pb.UnimplementedChatServer
	pool []pb.Chat_BroadcastMessageServer
}

func (c *Chat) BroadcastMessage(stream pb.Chat_BroadcastMessageServer) error {
	c.pool = append(c.pool, stream)

	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		for i := range c.pool {
			if err := c.pool[i].Send(msg); err != nil {
				log.Printf("Failed to send message: %v", err)
			}
		}
	}
}

func main() {
	// Initialize the gRPC server
	grpcServer := grpc.NewServer()

	// Register the Chat service
	pb.RegisterChatServer(grpcServer, &Chat{
		pool: make([]pb.Chat_BroadcastMessageServer, 0, 10),
	})

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50051...")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
