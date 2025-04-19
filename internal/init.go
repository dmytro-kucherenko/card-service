package internal

import (
	restApi "github.com/dmytro-kucherenko/card-service/internal/api/rest"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/config"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/multiplexer"
	"github.com/gofiber/fiber/v2"
)

var fiberApp *fiber.App

func init() {
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
		WithFiber(fiberApp).
		ServeGracefully()

	if err != nil {
		logger.Fatal(err)
	}
}
