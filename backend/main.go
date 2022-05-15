package main

import (
	"context"
	"errors"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"

	"github.com/tk42/victolinux/env"
	"github.com/tk42/victolinux/logging/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ADDRESS = "0.0.0.0:8080"
)

func authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	if token != env.GetString("GRPC_AUTH_TOKEN", "") {
		return nil, errors.New("unauthorized")
	}
	return ctx, nil
}

func main() {
	c := context.Background()

	logger := logging.GetLogger("umlquiz", logging.LogLevel(zapcore.DebugLevel))
	logger.Info("sever start", zap.String("address", ADDRESS))

	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(err)
	}

	var s *grpc.Server
	var isAuthSet bool
	if env.GetString("GRPC_AUTH_TOKEN", "") == "" {
		s = grpc.NewServer()
		isAuthSet = false
	} else {
		authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticate))
		s = grpc.NewServer(authInterceptor)
		isAuthSet = true
	}

	presentation := InjectPresentation()
	umlquiz.RegisterUMLQuizServiceServer(s, &presentation)
	reflection.Register(s)

	logger.Info("starting to listen", zap.Bool("isAuthSet", isAuthSet))

	go func(ctx context.Context) {
		<-ctx.Done()
		s.GracefulStop()
	}(c)

	if err := s.Serve(listener); err != nil {
		defer c.Done()
		panic(err)
	}
}
