package multiplexer

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/soheilhy/cmux"
)

type httpAdapter struct {
	app *fiber.App
}

func newFiberAdapter(app *fiber.App) Server {
	return &httpAdapter{app}
}

func (adapter *httpAdapter) Serve(server cmux.CMux) error {
	listener := server.Match(cmux.HTTP1Fast())

	err := adapter.app.Listener(listener)
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (adapter *httpAdapter) Stop() {
	adapter.app.Shutdown()
}
