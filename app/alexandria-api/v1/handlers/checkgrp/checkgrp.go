package checkgrp

import (
	"context"
	"net/http"
	"os"
	"runtime"

	"github.com/charlieroth/alexandria/foundation/web"
)

type handlers struct{}

func new() *handlers {
	return &handlers{}
}

// liveness returns a simple status info if the service is alive
// TODO: Add more information in a k8s environment
func (h *handlers) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status     string `json:"status,omitempty"`
		Host       string `json:"host,omitempty"`
		GOMAXPROCS int    `json:"GOMAXPROCS,omitempty"`
	}{
		Status:     "up",
		Host:       host,
		GOMAXPROCS: runtime.GOMAXPROCS(0),
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}
