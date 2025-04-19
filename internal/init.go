package internal

import (
	grpcApi "github.com/dmytro-kucherenko/card-service/internal/api/grpc"
	restApi "github.com/dmytro-kucherenko/card-service/internal/api/rest"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/config"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/multiplexer"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

var (
	grpcServer *grpc.Server
	fiberApp   *fiber.App
)

func init() {
	grpcServer = grpcApi.Run()
	fiberApp = restApi.Run()
}

func Run() {
	logger := log.NewConsole("Server")
	port := config.AppPort()
	instance, err := multiplexer.New(port)

	if err != nil {
		logger.Fatal(err)
	}

	err = instance.
		WithLogger(logger).
		WithGRPC(grpcServer).
		WithFiber(fiberApp).
		ServeGracefully()

	if err != nil {
		logger.Fatal(err)
	}
}
