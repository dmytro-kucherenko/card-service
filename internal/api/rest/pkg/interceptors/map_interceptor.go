package interceptors

import (
	"github.com/gofiber/fiber/v2"
)

func Map(mapError func(error) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		return mapError(err)
	}
}
