package register

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type Router = fiber.Router

func Get[P, R any](router Router, path string, status int, handle func(context.Context, P) (R, error)) Router {
	return router.Get(path, route(status, handle, false))
}

func Post[P, R any](router Router, path string, status int, handle func(context.Context, P) (R, error)) Router {
	return router.Post(path, route(status, handle, true))
}

func Patch[P, R any](router Router, path string, status int, handle func(context.Context, P) (R, error)) Router {
	return router.Patch(path, route(status, handle, true))
}

func Delete[P, R any](router Router, path string, status int, handle func(context.Context, P) (R, error)) Router {
	return router.Delete(path, route(status, handle, false))
}
