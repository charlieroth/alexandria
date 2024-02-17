package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"syscall"
)

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	mux      *http.ServeMux
	shutdown chan os.Signal
}

// NewApp creates an App value that handles a set of routes for the application
func NewApp(shutdown chan os.Signal) *App {
	mux := http.NewServeMux()
	return &App{
		mux:      mux,
		shutdown: shutdown,
	}
}

// SignalShutdown is used to gracefully shutdown the application when
// and integrity issue is identified
func (app *App) SignalShutdown() {
	app.shutdown <- syscall.SIGINT
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.mux.ServeHTTP(w, r)
}

func (app *App) Handle(method string, group string, path string, handler Handler) {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		if err := handler(r.Context(), w, r); err != nil {
			if validateError(err) {
				app.SignalShutdown()
				return
			}
		}
	}

	fullPath := path
	if group != "" {
		fullPath = "/" + group + path
	}

	fullPath = fmt.Sprintf("%s %s", method, fullPath)
	app.mux.HandleFunc(fullPath, handlerFunc)
}

// validateError validates the error for special conditions that do not warrant an actual
// shutdown by the system
func validateError(err error) bool {
	if errors.Is(err, syscall.EPIPE) {
		return false
	}

	if errors.Is(err, syscall.ECONNRESET) {
		return false
	}

	return true
}
