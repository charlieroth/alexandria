package all

import (
	"github.com/charlieroth/alexandria/app/alexandria-api/v1/handlers/checkgrp"
	"github.com/charlieroth/alexandria/business/web/v1/mux"
	"github.com/charlieroth/alexandria/foundation/web"
)

// Routes constructs the add value which provides the implementation
// of RouteAdder for specifying what route to bind to this instance
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouteAdder interface
func (add) Add(app *web.App, muxConfig mux.Config) {
	checkgrp.Routes(app, checkgrp.Config{})
}
