package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charlieroth/alexandria/app/alexandria-api/v1/build/all"
	"github.com/charlieroth/alexandria/business/web/v1/mux"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// TODO: create configuration

	// TODO: Open database connection

	// TODO: Load private keys from disk, usually injected by some system like Vault

	// Start service
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	muxConfig := mux.Config{
		Shutdown: shutdown,
	}

	api := http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      mux.WebApi(muxConfig, buildRoutes()),
		ReadTimeout:  time.Duration(5 * time.Second),
		WriteTimeout: time.Duration(5 * time.Second),
		IdleTimeout:  time.Duration(120 * time.Second),
	}

	serverErrors := make(chan error, 1)

	go func() {
		fmt.Println("listening on http://0.0.0.0:3000")
		serverErrors <- api.ListenAndServe()
	}()

	// Shutdown
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-shutdown:
		log.Default()
		ctx, cancel := context.WithTimeout(ctx, time.Duration(20*time.Second))
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

func buildRoutes() mux.RouteAdder {
	return all.Routes()
}
