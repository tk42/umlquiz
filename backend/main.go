package main

import (
	"net"

	"github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/umlquiz/backend/utility"

	"github.com/tk42/victolinux/logging/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	ADDRESS = "0.0.0.0:8080"
)

func main() {
	logger := logging.GetLogger("umlquiz", logging.LogLevel(zapcore.DebugLevel))
	logger.Info("sever start")

	logger.Info("starting to listen", zap.String("address", ADDRESS))
	listener, err := net.Listen("tcp", ADDRESS)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	presentation := utility.InjectPresentation()
	umlquiz.RegisterUMLQuizServiceServer(server, &presentation)

	if err := server.Serve(listener); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
