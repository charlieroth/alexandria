package checkgrp

import (
	"net/http"

	"github.com/charlieroth/alexandria/foundation/web"
)

// Config contains all the mandatory systems required for hanlders
// TODO: Build info, logger, DB handler
type Config struct{}

func Routes(app *web.App, config Config) {
	const version = "v1"
	handlers := new()
	app.Handle(http.MethodGet, version, "/liveness", handlers.liveness)
}
