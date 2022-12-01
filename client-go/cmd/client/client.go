package client

import (
	"context"
	"interview-client/internal/api/interview"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start() {
	serviceAddress := "127.0.0.1:8080"
	log.Printf("Connecting to interview service at %s", serviceAddress)

	conn, err := grpc.DialContext(
		context.Background(),
		serviceAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to connect to service"))
	}

	client := interview.NewInterviewServiceClient(conn)

	resp, err := client.HelloWorld(context.Background(), &interview.HelloWorldRequest{})
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to hello world"))
	}

	log.Println(resp)
}
