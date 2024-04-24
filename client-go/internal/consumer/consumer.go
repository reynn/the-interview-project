package consumer

import (
	"context"
	"fmt"
	"interview-client/internal/api/interview"
	"log/slog"

	"google.golang.org/grpc"
)

type consumer struct {
	interview.UnimplementedInterviewServiceServer
	client interview.InterviewServiceClient
}

func New(c *grpc.ClientConn) *consumer {
	return &consumer{
		client: interview.NewInterviewServiceClient(c),
	}
}

func (s *consumer) HelloWorld(ctx context.Context) {
	resp, err := s.client.HelloWorld(context.Background(), &interview.HelloWorldRequest{})
	if err != nil {
		slog.Error("failed to hello world", slog.Any("error", err))
	}
	fmt.Println(resp)
}
