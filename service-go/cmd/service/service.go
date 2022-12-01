package service

import (
	"fmt"
	"log"
	"net"

	"interview-service/internal/api"
	"interview-service/internal/api/interview"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {
	address := fmt.Sprintf("localhost:%d", 8080)
	log.Printf("Starting interview service at %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}
