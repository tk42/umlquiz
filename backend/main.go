package main

import (
	"context"
	"errors"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/umlquiz/backend/utility"

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
	logger := logging.GetLogger("umlquiz", logging.LogLevel(zapcore.DebugLevel))
	logger.Info("sever start", zap.String("address", ADDRESS))

	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(err)
	}

	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticate))
	server := grpc.NewServer(authInterceptor)
	reflection.Register(server)

	presentation := utility.InjectPresentation()
	umlquiz.RegisterUMLQuizServiceServer(server, &presentation)

	logger.Info("starting to listen")
	if err := server.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
