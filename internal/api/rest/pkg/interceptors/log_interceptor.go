package interceptors

import (
	"fmt"
	"time"

	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/gofiber/fiber/v2"
)

func Logger(logger log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		end := time.Since(start)

		latency := float64(end.Microseconds()) / 1000.0
		code := c.Response().StatusCode()
		path := c.Route().Path
		message := fmt.Sprintf("[%v] %v %vms", code, path, latency)

		if err == nil {
			logger.Info(message)
		} else {
			logger.Warn(fmt.Sprintf("%v\n%v", message, err.Error()))
		}

		return err
	}
}
