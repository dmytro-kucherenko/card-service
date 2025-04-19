package grpc

import (
	"github.com/dmytro-kucherenko/card-service/internal/api/grpc/card"
	"github.com/dmytro-kucherenko/card-service/internal/api/grpc/pkg/interceptors"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() *grpc.Server {
	logger := log.NewConsole("GRPC")
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptors.LogUnary(logger),
		interceptors.ErrorUnary(logger),
		interceptors.ValidateUnary()),
	)

	card.NewHandler().Init(server)
	reflection.Register(server)

	logger.Info("Server created")

	return server
}
