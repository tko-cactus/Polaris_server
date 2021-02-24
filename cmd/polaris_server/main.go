package main

import (
	"context"
	"errors"
	"log"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":51411")
	if err != nil {
		log.Fatal(err)
	}
	authInterceptor := grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(authenticate))
	server := grpc.NewServer(authInterceptor)
	server.Serve(listenPort)
}

func authenticate(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	if token != "testtoken" {
		return nil, errors.New("unauthorized")
	}
	return ctx, nil
}
