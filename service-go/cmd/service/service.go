package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"

	config "interview-service/config"
	"interview-service/internal/api"
	"interview-service/internal/api/interview"
	jwt "interview-service/internal/domain/jwt"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	appCfg, loadErr := config.Load()
	if loadErr != nil {
		slog.Error("failed to load the app config", slog.Any("error", loadErr))
		os.Exit(1)
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: func()slog.Leveler{
			if appCfg.Debug {
				return slog.LevelDebug
			}
			return slog.LevelInfo
		}(),
	})))

	address := fmt.Sprintf("%s:%s", appCfg.GRPC.ServerHost, appCfg.GRPC.UnsecurePort)

	lis, listenErr := net.Listen("tcp", address)
	if listenErr != nil {
		slog.Error("failed to listen", slog.Any("error", listenErr), slog.String("address", address))
		os.Exit(1)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(validateJWT([]byte(appCfg.JWTSecret))),
		),
	}

	grpcServer := grpc.NewServer(opts...)

	interview.RegisterInterviewServiceServer(grpcServer, api.New())
	reflection.Register(grpcServer)

	slog.Info("Starting interview service", slog.String("address", address))
	grpcServer.Serve(lis)
}

const (
	authHeader = "authorization"
)

// validateJWT parses and validates a bearer jwt
//
// TODO: move to own package (in ./internal/api/auth) using a constructor that privately sets the secret
func validateJWT(secret []byte) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		claims, err := jwt.ValidateToken(token, secret)
		if err != nil {
			slog.Error("error validating jwt token", slog.Any("error", err))
			return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		ctx = context.WithValue(ctx, authHeader, claims)

		return ctx, nil
	}
}
