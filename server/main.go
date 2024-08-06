package main

import (
	"log"
	"net"
	"src/pb"
	"src/server/controllers"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")

	server := controllers.NewServer()

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductsServiceServer(s, server)

	log.Println("Server is running on port :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
