package main

import (
	"context"
	"interview-client/internal/config"
	"interview-client/internal/consumer"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	cfg, cfgErr := config.NewFromEnv()
	if cfgErr != nil {
		slog.Error("failed to retrieve config from environment", slog.Any("error", cfgErr))
		os.Exit(1)
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: func() slog.Leveler {
			if cfg.Debug {
				return slog.LevelDebug
			}
			return slog.LevelInfo
		}(),
	})))

	conn, dialErr := grpc.DialContext(
		ctx,
		cfg.ServerAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if dialErr != nil {
		slog.Error("failed to connect to service", slog.Any("error", dialErr))
		os.Exit(1)
	}

	consumer := consumer.New(conn)

	consumer.HelloWorld(ctx)
}
