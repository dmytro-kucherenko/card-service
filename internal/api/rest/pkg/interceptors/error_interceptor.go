package interceptors

import (
	"net/http"

	restErrors "github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/errors"
	"github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/models"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/gofiber/fiber/v2"
)

func sendError(c *fiber.Ctx, code appErrors.ErrCode, status int, message string) error {
	return c.Status(status).JSON(models.ErrorResponse{
		Response: models.Response{Valid: false},
		Error: models.Error{
			Code:    code.String(),
			Message: message,
		},
	})
}

func Error(logger log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if errPanic := recover(); errPanic != nil {
				logger.Error("Panic recovery: ", errPanic)
				err = sendError(c, appErrors.ErrInternal, http.StatusInternalServerError, "internal server error")
			}
		}()

		err = c.Next()
		if err == nil {
			return nil
		}

		code, okCode := appErrors.Code(err)
		status, okStatus := restErrors.Status(err)

		if !okCode || !okStatus {
			logger.Error(err.Error())

			return sendError(c, appErrors.ErrInternal, http.StatusInternalServerError, "internal server error")
		}

		return sendError(c, code, status, err.Error())
	}
}
