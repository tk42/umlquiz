package main

import (
	"context"
	"fmt"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpczap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	auth "github.com/tk42/jwt-go-auth"
	"github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/victolinux/env"

	"github.com/tk42/victolinux/logging/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := context.Background()

	address := fmt.Sprintf("0.0.0.0:%s", env.GetString("PORT", "8080"))

	logger := logging.GetLogger("umlquiz", logging.LogLevel(zapcore.DebugLevel))
	logger.Info("sever start", zap.String("address", address))

	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	loginServer := InjectServerUnauthenticated()

	env := auth.Env{SecretSigningKey: loginServer.SigningSecret}
	authInterceptor := grpc.UnaryInterceptor(
		grpcMiddleware.ChainUnaryServer(
			grpczap.UnaryServerInterceptor(logger.Logger),
			grpc_auth.UnaryServerInterceptor(env.AuthFunc),
		),
	)
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()), authInterceptor)
	// s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// s := grpc.NewServer(grpc.WithInsecure()) // Deprecated
	umlquiz.RegisterUMLQuizLoginServiceServer(s, &loginServer)

	presentation := InjectPresentation()
	umlquiz.RegisterUMLQuizHelloServiceServer(s, &presentation)
	umlquiz.RegisterUMLQuizUserServiceServer(s, &presentation)
	umlquiz.RegisterUMLQuizQuizServiceServer(s, &presentation)
	umlquiz.RegisterUMLQuizReportServiceServer(s, &presentation)
	reflection.Register(s)

	logger.Info("starting to listen")

	go func(ctx context.Context) {
		<-ctx.Done()
		s.GracefulStop()
	}(c)

	if err := s.Serve(listener); err != nil {
		defer c.Done()
		panic(err)
	}
}
