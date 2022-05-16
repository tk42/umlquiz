package adapter

import (
	"bytes"
	"context"

	auth "github.com/tk42/jwt-go-auth"
	autogen "github.com/tk42/umlquiz/backend/gen/proto/golang/github.com/tk42/umlquiz"
	"github.com/tk42/victolinux/env"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerUnauthenticated struct {
	autogen.UnimplementedUMLQuizLoginServiceServer
	SigningSecret []byte
}

func NewServerUnauthenticated() ServerUnauthenticated {
	secretBytes := []byte(env.GetString("GRPC_AUTH_TOKEN", ""))
	if bytes.Equal(secretBytes, []byte("")) {
		panic("GRPC_AUTH_TOKEN is empty")
	}
	return ServerUnauthenticated{
		SigningSecret: secretBytes,
	}
}

// AuthFuncOverride is called instead of AuthFunc for methods on ServerUnauthenticated struct
func (s *ServerUnauthenticated) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	// fmt.Println("client is calling method:", fullMethodName)

	////this check is optional as it should no be possible to get here anyway
	// if fullMethodName != "/umlquiz.UMLQuizLoginService/GetToken" {
	// 	return nil, status.Errorf(codes.Unauthenticated, "no token")
	// }

	return ctx, nil
}

func (s *ServerUnauthenticated) GetToken(ctx context.Context, req *autogen.GetTokenRequest) (*autogen.GetTokenResponse, error) {
	tokenSigned, err := auth.GenerateToken(req.Username, s.SigningSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not generate token")
	}
	return &autogen.GetTokenResponse{Token: tokenSigned}, nil
}
