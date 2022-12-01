package api

import (
	"context"

	"interview-service/internal/api/interview"
	"interview-service/internal/domain/greeter"
)

type server struct {
	interview.UnimplementedInterviewServiceServer
}

func New() *server {
	return &server{}
}

func (s *server) HelloWorld(ctx context.Context, r *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	greeting := greeter.Greet(r.GetName())

	return &interview.HelloWorldResponse{
		Greeting: greeting,
	}, nil
}
