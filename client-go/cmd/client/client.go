package main

import (
	"context"
	"interview-client/internal/config"
	"interview-client/internal/consumer"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	cfg, cfgErr := config.NewFromEnv()
	if cfgErr != nil {
		log.Fatalf("failed to retrieve config from environment: %v", cfgErr)
	}

	conn, dialErr := grpc.DialContext(
		ctx,
		cfg.ServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if dialErr != nil {
		log.Fatalf("failed to connect to service: %v", dialErr)
	}

	consumer := consumer.New(conn)

	consumer.HelloWorld(ctx)
}
