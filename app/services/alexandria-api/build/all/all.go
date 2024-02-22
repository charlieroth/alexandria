package all

import (
	"github.com/charlieroth/alexandria/app/services/alexandria-api/handlers/checkgrp"
	"github.com/charlieroth/alexandria/business/web/mux"
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
