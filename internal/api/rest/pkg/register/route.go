package register

import (
	"context"
	"net/http"
	"reflect"

	"github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/errors"
	appErrors "github.com/dmytro-kucherenko/card-service/internal/pkg/errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func route[P, R any](status int, handle func(context.Context, P) (R, error), bindBody bool) fiber.Handler {
	validate := validator.New()

	return func(c *fiber.Ctx) error {
		var params P

		if paramsType := reflect.TypeOf(params); paramsType != nil {
			steps := []struct {
				handle   func(any) error
				validate bool
			}{
				{c.BodyParser, bindBody},
				{c.ParamsParser, true},
				{c.QueryParser, true},
				{validate.Struct, true},
			}

			for _, step := range steps {
				if step.validate {
					if err := step.handle(&params); err != nil {
						return errors.NewHTTPError(http.StatusBadRequest, appErrors.ErrValidation, err.Error())
					}
				}
			}
		}

		result, err := handle(c.Context(), params)
		if err != nil {
			return err
		}

		return c.Status(status).JSON(result)
	}
}
