package rest

import (
	"github.com/dmytro-kucherenko/card-service/api/openapi"
	"github.com/dmytro-kucherenko/card-service/internal/api/rest/card"
	"github.com/dmytro-kucherenko/card-service/internal/api/rest/pkg/interceptors"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/config"
	"github.com/dmytro-kucherenko/card-service/internal/pkg/log"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

func init() {
	host := config.AppHost()
	basePath := config.AppBasePath()
	protocol := config.AppProtocol()

	openapi.SwaggerInfo.Title = "Card API"
	openapi.SwaggerInfo.Description = "API server for processing bank card requests."
	openapi.SwaggerInfo.Version = "1.0"
	openapi.SwaggerInfo.Host = host
	openapi.SwaggerInfo.BasePath = basePath
	openapi.SwaggerInfo.Schemes = []string{protocol}
}

// @version		1.0
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func Run() *fiber.App {
	logger := log.NewConsole("REST")
	basePath := config.AppBasePath()

	app := fiber.New()
	api := app.Group(basePath).Use(
		interceptors.Logger(logger),
		interceptors.Error(logger),
	)

	app.Get("/swagger/*", swagger.HandlerDefault)

	card.NewHandler().Init(api)

	logger.Info("Server created")

	return app
}
