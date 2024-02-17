// Provides support to bind domain level routes to application mux
package mux

import (
	"net/http"
	"os"

	"github.com/charlieroth/alexandria/foundation/web"
)

type Config struct {
	Shutdown chan os.Signal
}

type RouteAdder interface {
	Add(app *web.App, config Config)
}

func WebApi(config Config, routeAdder RouteAdder) http.Handler {
	app := web.NewApp(config.Shutdown)
	routeAdder.Add(app, config)
	return app
}
